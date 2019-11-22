package main

import (
	"fmt"
	"io"
	"os"
)

type IO interface {
	io.ReaderAt
	io.WriterAt
	io.Seeker
	Fd() uintptr
}

var quit = make(chan struct{})

func open(fd uintptr, name string) {
	file := os.NewFile(fd, "")
	_, err := file.Write([]byte(`testing is in progress`))
	if err != nil {
		panic("write failed")
	}
	txt := make([]byte, 100)
	n, err := file.ReadAt(txt, 0)
	if err != nil {
		fmt.Println(err, n)
	}
	fmt.Println(string(txt))
	fmt.Println("FD2: %v", file.Fd())
	quit <- struct{}{}
}

func main() {
	var file []IO
	fmt.Println("hello: ", file)
	file = []IO{nil}
	name := "my name is utkarsh"
	fmt.Println("hello: ", file)
	f, _ := os.OpenFile("hello.txt", os.O_RDWR, 06666)
	file = append(file, f)
	fmt.Println(file)
	n, err := file[1].WriteAt([]byte(name), 50)
	if err != nil {
		fmt.Println(err, n)
	}
	fmt.Println(n)
	txt := make([]byte, 100)
	n, err = file[1].ReadAt(txt, 50)
	if err != nil {
		fmt.Println(err, n)
	}
	fmt.Println("FD1: %v", file[1].Fd())

	go open(file[1].Fd(), name)
	fmt.Println(n)
	fmt.Println(string(txt))
	<-quit
}
