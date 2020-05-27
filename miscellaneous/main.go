package main

import "fmt"

func main() {
	var val int
	fmt.Println(&val)
	f(10000)
	fmt.Println(&val)
}

func f(i int) {
	if i--; i == 0 {
		return
	}
	f(i)
}
