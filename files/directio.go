package main

import (
	"fmt"
	"os"
	"syscall"

	"github.com/ncw/directio"
)

/* a&(b-1) = a%b
 */

func main() {
	fmt.Println("opening file")
	file, err := os.OpenFile("file1.txt", os.O_TRUNC|syscall.O_DIRECT|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}

	buf := []byte("Hi! There")
	fmt.Println(&buf[0])
	newbuf := directio.AlignedBlock(len(buf))
	copy(newbuf, buf)
	_, err = file.Write(buf)
	if err != nil {
		panic(err)
	}
	fmt.Println("Removing")
	if err := os.Remove("file1.txt"); err != nil {
		panic(err)
	}

	fmt.Println("Closing")
	if err := file.Close(); err != nil {
		panic(err)
	}
}
