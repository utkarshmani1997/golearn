package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	val := os.Args[1]
	for {
		switch val {
		case "1":
			fmt.Println("1")
			if val == "1" {
				break
			}
			fmt.Println("2")
		}
		time.Sleep(1 * time.Second)
		fmt.Println("3")
	}
}
