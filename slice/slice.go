package main

import "fmt"

func change(a []int) {
	for i, _ := range a {
		a[i] += 10
	}
}

func main() {
	var a = []int{1, 2, 3, 4, 5, 7}
	change(a)
	fmt.Println(a[:], a[:5], a[2:5])
}
