package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/utkarshmani1997/golearn/json/types"
)

func main() {
	peter := types.User{
		Firstname: "John",
		Lastname:  "Doe",
	}
	person, _ := json.Marshal(peter)
	for {
	up:
		ip := os.Getenv("MY_SERVER_IP")
		fmt.Println(ip)
		resp, err := http.Post("http://"+ip+":8080/decode", "application/json", bytes.NewBuffer(person))
		if err != nil {
			fmt.Println("err: ", err)
			time.Sleep(5 * time.Second)
			goto up
		}
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
		time.Sleep(5 * time.Second)
	}
}
