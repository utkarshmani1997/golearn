// +build build3

package main

import "fmt"

var DEBUG = "False"

func main() {
	if DEBUG == "True" {
		fmt.Println("This is build3 with DEBUG = true")
	} else {
		fmt.Println("This is build3 with DEBUG = false")
	}
}
