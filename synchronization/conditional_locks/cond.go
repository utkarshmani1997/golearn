package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

var (
	count = 0
	m     = sync.Mutex{}
	c     = sync.NewCond(&m)
)

func main() {
	ch := make(chan bool)
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	go func() {
		for {
			fmt.Println("Wait goroutine started")
			c.L.Lock()
			fmt.Println("Waiting for broadcast")
			c.Wait()
			fmt.Println("Received broadcast, count: ", count)
			if count == 5 {
				fmt.Println("It's from goroutine that starts a task")
				start()
				c.L.Unlock()
			}
		}
	}()

	time.Sleep(1 * time.Second)
	counter(ch)
	//	countagain(ch)
	<-ch
	time.Sleep(1 * time.Minute)
}

func counter(ch chan bool) {
	go func(ch chan bool) {
		c.L.Lock()
		defer c.L.Unlock()
		for {
			fmt.Println("count: ", count)
			if count == 5 {
				c.Broadcast()
				time.Sleep(2 * time.Second)
				ch <- true
				return
			}
			count++
		}
	}(ch)
}

func countagain(ch chan bool) {
	for <-ch {
		time.Sleep(10 * time.Millisecond)
		count = 0
		counter(ch)
	}
}

func start() {
	cnt := 0
	for {
		if cnt == 5 {
			fmt.Println("start")
			time.Sleep(10 * time.Millisecond)
			cnt++
			return
		}
		cnt++
	}
}
