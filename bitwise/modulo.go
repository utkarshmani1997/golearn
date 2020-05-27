package main

import "fmt"

const diviser = 4096

func modulo(value int) int {
	return value % diviser
}

func moduloBitwise(value int) int {
	return value & (diviser - 1)
}

func main() {
	fmt.Println(modulo(13010818623874))
	fmt.Println(moduloBitwise(1301081))
}
