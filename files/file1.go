package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("opening file")
	file, err := os.OpenFile("file1.txt", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Get fd")
	fd := file.Fd()
	if fd < 0 {
		panic("fd not opened")
	}
	/*	if err := syscall.Flock(int(fd), syscall.LOCK_EX|syscall.LOCK_NB); err != nil {
			panic(err)
		}
		for {
		}
	*/
	fmt.Println("Removing")
	if err := os.Remove("file1.txt"); err != nil {
		panic(err)
	}

	fmt.Println("Closing")
	if err := file.Close(); err != nil {
		panic(err)
	}
}
