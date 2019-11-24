package main

import "fmt"

func main() {
	var a = []int{4, 1, 7, 2, 12, 9, 10, 13, 5, 3, 6, 8}
	for i := 1; i < len(a); i++ {
		key := a[i]
		j := i - 1
		for j >= 0 && a[j] > key {
			a[j+1] = a[j]
			j = j - 1
		}
		a[j+1] = key
	}
	fmt.Println(a)
}
