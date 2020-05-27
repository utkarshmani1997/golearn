package main

import (
	"fmt"
	"io"
	"os"
	"strings"
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

type x struct {
	x int64
	y int64
}

func main() {
	var file []IO
	//	fmt.Println("hello: ", file)
	file = []IO{nil}
	//	name := "my name is utkarsh"
	//y := x{1, 2}
	//	fmt.Println("hello: ", file)
	f, err := os.OpenFile("hello.txt", os.O_RDWR, 06666)
	if err != nil {
		panic(err)
	}
	file = append(file, f)
	//	fmt.Println(file)
	// buf := make([]byte, 4096)
	//	var buf bytes.Buffer
	//	//str := strconv.FormatInt(100000, 10) + "," + name
	//	//copy(buf, []byte(str))
	//	err = binary.Write(&buf, binary.BigEndian, y)
	//	fmt.Println(string(buf.Bytes()), err)
	//	n, err := file[1].WriteAt(buf.Bytes(), 0)
	//	if err != nil {
	//		panic(err)
	//	}
	var n int
	txt := make([]byte, 40)
	n, err = file[1].ReadAt(txt, 0)
	if err != nil {
		fmt.Println(err, n)
		panic("read fail")
	}
	println(n)
	//	ioutil.ReadAll()
	nam := strings.Trim(string(txt), "\x00")
	na := strings.Split(nam, ",")
	fmt.Println(na, len(na), string(txt))
	/*
		fmt.Println("FD1: %v", file[1].Fd())

		go open(file[1].Fd(), name)
		fmt.Println(n)
		fmt.Println(string(txt))
		<-quit
	*/
}
