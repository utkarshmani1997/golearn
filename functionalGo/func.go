package main

import "fmt"

type Operation func(x, y int) int

// f(x, y) = x + y
func Add() Operation {
	return func(x, y int) int {
		return x + y
	}
}

// f(x, y) = x - y
func Substract() Operation {
	return func(x, y int) int {
		return x - y
	}
}

// f(x, y) = x * y
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

func operation(f func() Operation, x, y int) {
	fmt.Printf("%d\n", f()(x, y))
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

	// this is readable, yes offcourse it's bit difficult
	// to understand if you are new to functional code.
	operation(Add, 10, 20)
	operation(Substract, 10, 5)
	operation(Multiply, 5, 5)
}
