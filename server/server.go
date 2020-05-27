package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, the end point is :%s!", r.URL.Path[1:])
}

func ReadHandler(w http.ResponseWriter, r *http.Request) {
	dat, err := ioutil.ReadFile("data.txt")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Content in file is...\r\n%s", string(dat))
	time.Sleep(5 * time.Second)
	panic("i'm dying")
}

func main() {
	http.HandleFunc("/", Handler)
	http.HandleFunc("/read", ReadHandler)
	http.ListenAndServe(":8080", nil)
}
