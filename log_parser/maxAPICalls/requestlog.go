package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

type LogEntry struct {
	Method     string
	URI        string
	StatusCode string
}

type LogEntryDetails struct {
	LogEntry
	Count int64
	Avg   float64
	Max   float64
	Min   float64
}

type Duration struct {
	Max   float64
	Min   float64
	Avg   float64
	Total float64
	Count int64
}

func main() {
	var logFiles []string

	// Find all xray-request.log files
	err := filepath.Walk(os.Args[1], func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.Name() == "xray-request.log" {
			logFiles = append(logFiles, path)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path %s: %v\n", os.Args[0], err)
		return
	}
	logCount := make(map[LogEntry]Duration)
	var startTime time.Time
	var endTime time.Time
	for _, logFile := range logFiles {
		file, err := os.Open(logFile)
		if err != nil {
			fmt.Printf("Error opening file: %v\n", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text()
			parts := strings.Split(line, "|")
			if len(parts) != 11 {
				fmt.Printf("Skipping malformed line: %s\n", line)
				continue
			}

			t, _ := time.Parse(time.RFC3339, parts[0])
			if startTime.IsZero() {
				startTime = t
			}
			if t.Before(startTime) {
				startTime = t
			}
			method := parts[4]
			uri := parts[5]
			statusCode := parts[6]
			duration, _ := strconv.ParseFloat(parts[9], 64)
			entry := LogEntry{
				Method:     method,
				URI:        uri,
				StatusCode: statusCode,
			}

			lc, ok := logCount[entry]
			if ok {
				logCount[entry] = Duration{
					Max: func() float64 {
						if lc.Max > duration {
							return lc.Max
						}
						return duration
					}(),
					Min: func() float64 {
						if lc.Min != float64(0) && lc.Min < duration {
							return lc.Min
						}
						return duration
					}(),
					Total: duration + lc.Total,
					Count: logCount[entry].Count + 1,
				}
			} else {
				logCount[entry] = Duration{
					Max:   0,
					Min:   0,
					Avg:   0,
					Count: 1,
				}
			}
			if endTime.IsZero() || t.After(endTime) {
				endTime = t
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("Error reading file: %v\n", err)
		}
	}
	var logEntries []LogEntryDetails
	for entry, d := range logCount {
		logEntries = append(logEntries, LogEntryDetails{
			LogEntry: entry,
			Count:    d.Count,
			Avg:      float64(d.Total / float64(d.Count)),
			Max:      d.Max,
			Min:      d.Min,
		})
	}

	// Sort log entries by count in descending order
	sort.Slice(logEntries, func(i, j int) bool {
		return logEntries[i].Count > logEntries[j].Count
	})
	// Print the grouped log counts
	cnt := 0
	fmt.Printf("Start Time: %s\n", startTime)
	for _, entry := range logEntries {
		if cnt == 10 {
			break
		}
		cnt++
		fmt.Printf("Method: %s, URI: %s, Status Code: %s, Count: %d, Min: %f, Avg: %f, Max: %f\n",
			entry.Method, entry.URI, entry.StatusCode, entry.Count, entry.Min, entry.Avg, entry.Max)
	}
	fmt.Printf("End time: %s\n", endTime)
}
