package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var logFiles []string

	// Find all xray-request.log files
	/*	err := filepath.Walk(os.Args[1], func(path string, info os.FileInfo, err error) error {
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
		logCount := make(map[string]time.Duration)
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

				trace_id := parts[1]
				duration, _ := strconv.ParseFloat(parts[9], 64)
				if duration > 1000000 {
					logCount[trace_id] = time.Duration(duration)
				}
			}

			if err := scanner.Err(); err != nil {
				fmt.Printf("Error reading file: %v\n", err)
			}
		}

	*/
	err := filepath.Walk(os.Args[1], func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.Contains(info.Name(), "xray-server-service") {
			logFiles = append(logFiles, path)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path %s: %v\n", os.Args[0], err)
		return
	}
	cnt := 0
	for traceID, entry := range []string{"01c8f3713098b767"} {
		if cnt == 1 {
			break
		}
		cnt++
		for _, logFile := range logFiles {
			file, err := os.Open(logFile)
			if err != nil {
				fmt.Printf("Error opening file: %v\n", err)
				return
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)

			for scanner.Scan() {
				log := scanner.Text()
				if strings.Contains(log, entry) {
					fmt.Println(traceID, logFile)
					fmt.Println(log)
				}
			}
			if err := scanner.Err(); err != nil {
				fmt.Printf("Error reading file: %v\n", err)
			}
		}
	}
}
