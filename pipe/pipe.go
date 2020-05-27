package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"time"
)

func main() {
	client, server := net.Pipe()
	go func() {
		for {
			time.Sleep(2 * time.Second)
			fmt.Println("Read at server")
			b, err := ioutil.ReadAll(server)
			if err != nil {
				panic(err)
			}
			fmt.Println("Client: ", string(b))
			_, err = server.Write([]byte("Hello, how are you?"))
			if err != nil {
				panic(err)
			}
		}
	}()

	time.Sleep(1 * time.Second)
	for {
		fmt.Println("Write at client")
		_, err := client.Write([]byte("Hi there!"))
		if err != nil {
			panic(err)
		}
		client.Close()
		client, server = net.Pipe()
		fmt.Println("Read at client")
		b, err := ioutil.ReadAll(client)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(b))
		time.Sleep(1 * time.Second)
	}
}
