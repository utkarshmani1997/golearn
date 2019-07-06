package main

import "fmt"

type hello struct {
	something interface{}
	count     int
}

var (
	ch        = make(chan hello, 100000)
	quit      = make(chan struct{})
	completed = make(chan struct{})
	closeRead = make(chan struct{})
)

func main() {
	go read()
	go Close()
	go write()
	ch <- hello{}
	<-quit
	close(ch)
	close(quit)
	close(completed)
	close(closeRead)
	fmt.Println("Got quit signal")
}

func write() {
	h := hello{}
	for i := 0; i < 25000; i++ {
		h.count = i
		ch <- h
	}
	completed <- struct{}{}
}

func read() {
	for ch != nil || quit != nil {
		select {
		case c, ok := <-ch:
			//			if (hello{}) == c {
			//				fmt.Println("empty chan")
			//			}
			fmt.Println("it's something", c.something)
			if !ok {
				fmt.Println("ch is done")
				continue
			}
			fmt.Println(c)
		case _, ok := <-closeRead:
			if !ok {
				fmt.Println("closeRead is done")
				continue
			}
			fmt.Println("Got close Read signal")
			quit <- struct{}{}
			break
		}
	}
}

func Close() {
	<-completed
	fmt.Println("Got completed signal")
	closeRead <- struct{}{}
}
