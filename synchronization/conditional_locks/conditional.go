package main

import (
	"fmt"
	"sync"
	"time"
)

var sharedRsc = make(map[string]interface{})
var (
	m = sync.Mutex{}
	c = sync.NewCond(&m)
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		// this go routine wait for changes to the sharedRsc
		c.L.Lock()
		fmt.Println("May be i'm first")
		for len(sharedRsc) == 0 {
			c.Wait()
		}
		fmt.Println("goroutine1", sharedRsc["rsc1"])
		c.L.Unlock()
		wg.Done()
	}()

	go func() {
		// this go routine wait for changes to the sharedRsc
		c.L.Lock()
		fmt.Println("May be i'm second")
		for len(sharedRsc) == 0 {
			c.Wait()
		}
		fmt.Println("goroutine2", sharedRsc["rsc2"])
		c.L.Unlock()
		wg.Done()
	}()

	time.Sleep(2 * time.Second)
	// this one writes changes to sharedRsc
	c.L.Lock()
	fmt.Println("I've started")
	sharedRsc["rsc1"] = "foo"
	sharedRsc["rsc2"] = "bar"
	c.Broadcast()
	c.L.Unlock()
	wg.Wait()
}
