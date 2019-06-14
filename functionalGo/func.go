package main

import "fmt"

type Operation func(x, y int) int

func Add() Operation {
	return func(x, y int) int {
		return x + y
	}
}

func Substract() Operation {
	return func(x, y int) int {
		return x - y
	}
}

func Multiply() Operation {
	return func(x, y int) int {
		return x * y
	}
}

func DoManyOperations(ops ...Operation) {
	for _, o := range ops {
		fmt.Printf("%d\n", o(5, 6))
	}
}

func main() {
	fmt.Printf("%d\n", Add()(13, 14))
	fmt.Printf("%d\n", Substract()(15, 16))
	fmt.Printf("%d\n", Multiply()(17, 18))

	// i want to perform this task sequentially not right
	// now but later on given set of arguments.
	// What is the benifit ?
	// This makes code self documented
	operations := []Operation{Add(), Substract(), Multiply()}
	DoManyOperations(operations...)
}
