package main

import (
	"fmt"
	"time"
)

const (
	pingtimeout = 2 * time.Second
)

func ping(t chan struct{}) {
	timeout := func() <-chan time.Time {
		return time.After(pingtimeout)
	}()
	fmt.Println("waiting on timeout")
	time.Sleep(3 * time.Second)
	select {
	case time := <-timeout:
		fmt.Printf("timeout done, time: %#v\n", time)
		select {
		case t <- struct{}{}:
		default:
		}
	}

}

func main() {
	exit := make(chan struct{})
	timeout := make(chan struct{}, 5)
	go func(exit, timeout chan struct{}) {
		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-exit:
				fmt.Println("exiting...")
				return
			case <-ticker.C:
				fmt.Println("ping...")
				ping(timeout)
				fmt.Println("Done")

			}
		}
	}(exit, timeout)

	go func(t chan struct{}) {
		for {
			time.Sleep(3 * time.Second)
			select {
			case <-t:
			}
		}
	}(timeout)

	time.Sleep(10 * time.Second)
	close(timeout)
	exit <- struct{}{}
	fmt.Println("Hello, playground")
}
