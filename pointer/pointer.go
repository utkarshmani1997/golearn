package main

import (
	"fmt"
	"runtime"
)

type X struct {
	Y map[string]string
}

func main() {
	var x []X
	PrintMemUsage()
	for i := 0; i < 1000000; i++ {
		a := X{
			Y: map[string]string{"hi": "hello"},
		}
		x = append(x, a)
	}

	if x[0].Y != nil {

	}
	PrintMemUsage()
	y := x
	fmt.Println(len(x), len(y))
	for i := 0; i < 100; i++ {
		x = x[:len(x)-100]
	}
	fmt.Println(len(x), len(y))
	PrintMemUsage()
	runtime.GC()
	y = nil
	x[0].Y["hi"]
	PrintMemUsage()
	runtime.GC()
	PrintMemUsage()
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
