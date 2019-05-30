package main

import (
	"fmt"
	"time"
)

var (
	ch   = make(chan int, 10)
	quit = make(chan struct{})

	completed = make(chan struct{})
)

func main() {
	go read()
	go Close()
	go write()
	<-quit
	close(quit)
	fmt.Println("Got quit signal")
	time.Sleep(5 * time.Second)
}

func write() {
	for i := 0; i < 1000; i++ {
		ch <- i
	}
	completed <- struct{}{}
}

func read() {
	for ch != nil {
		time.Sleep(1 * time.Millisecond)
		c, ok := <-ch
		if !ok {
			fmt.Println("ch is done")
			quit <- struct{}{}
			break
		}
		fmt.Println(c)
	}
}

func Close() {
	<-completed
	fmt.Println("Completed")
	close(ch)
}
