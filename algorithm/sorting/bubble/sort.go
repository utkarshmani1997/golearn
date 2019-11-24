package main

import "fmt"

func main() {
	var a = []int{4, 1, 7, 2, 12, 9, 10, 13, 5, 3, 6, 8}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if a[j] < a[i] {
				tmp := a[j]
				a[j] = a[i]
				a[i] = tmp
			}
		}
	}
	fmt.Println(a)
}
