package main

import (
	"bufio"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os/exec"
	"strings"
)

type stats struct {
	poolName          string
	allocatedCapacity string
	freeCapacity      string
	reads             string
	writes            string
	readLatency       string
	writeLatency      string
}

func runCommand() {
	cmd := exec.Command("zpool", "iostat", "vol1", "1")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return
	}
	cmd.Start()
	r := bufio.NewReader(stdout)
	r.Discard(158)
	for true {
		str, _ := r.ReadString('\n')
		out := strings.Fields(str)
		stats := map[string]string{
			"poolName":          out[0],
			"allocatedCapacity": out[1],
			"freeCapacity":      out[2],
			"reads":             out[3],
			"writes":            out[4],
			"readLatency":       out[5],
			"writeLatency":      out[6],
		}
		fmt.Println("stats:", stats)
	}
}

func main() {
	go http.ListenAndServe("localhost:8080", nil)
	runCommand()
}
