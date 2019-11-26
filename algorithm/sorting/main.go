package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/utkarshmani1997/golearn/algorithm/sorting/bubble"
	"github.com/utkarshmani1997/golearn/algorithm/sorting/insertion"
	"github.com/utkarshmani1997/golearn/algorithm/sorting/merge"
)

const (
	availableFlags = `
Supporter flags: bubble, insertion, merge, mergesort
for exp: ./sort bubble
`
)

func generateRandom(size int) []int {
	arr := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(size)
	}
	return arr
}

func timed(f func(arr []int) []int, arr []int) []int {
	start := time.Now()
	result := f(arr)
	fmt.Println("Time to finish: ", time.Since(start))
	return result
}

func main() {
	size := flag.Int("size", 20, "Generate random array of given size to be sorted")
	flag.Parse()
	var a = generateRandom(*size)
	switch flag.Arg(0) {
	case "bubble":
		fmt.Println("Array to be sorted: ", a)
		fmt.Println("Sorted Array: ", timed(bubble.Sort, a))
	case "insertion":
		fmt.Println("Array to be sorted: ", a)
		fmt.Println("Sorted Array: ", timed(insertion.Sort, a))
	case "merge":
		a = []int{10, 20, 30, 40, 1, 5, 6, 9}
		fmt.Println("Array to be sorted: ", a)
		fmt.Println("Sorted Array: ", merge.Merge(a[0:len(a)/2], a[len(a)/2:]))
	case "mergesort":
		fmt.Println("Array to be sorted: ", a)
		fmt.Println("Sorted Array: ", timed(merge.Sort, a))
	default:
		fmt.Println("No flags entered")
		fmt.Println(availableFlags)
		os.Exit(1)
	}
}
