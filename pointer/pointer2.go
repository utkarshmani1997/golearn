package main

import (
	"fmt"
	"runtime"
)

type X struct {
	Y int
	s []string
	m map[int]string
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
func main() {
	PrintMemUsage()
	var x *X
	x = &X{
		Y: 5,
	}
	x.m = map[int]string{}
	for i := 0; i < 1000000; i++ {
		x.s = append(x.s, "hi")
		x.m[i] = "hi"
	}
	PrintMemUsage()
	x = nil
	//	runtime.GC()
	PrintMemUsage()
}
