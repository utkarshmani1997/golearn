package main

import "fmt"

func main() {
	x := func(x int) {
		fmt.Println(x)
	}
	x(10)
}
