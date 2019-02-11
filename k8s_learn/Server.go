package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, the end point is :  %s !", r.URL.Path[1:])
}

func readhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, the end point is :  %s !\r\n", r.URL.Path[1:])
	dat, err := ioutil.ReadFile("data.txt")
	if err != nil {
		fmt.Fprintf(w, "Error in reading file, error: %s ", err.Error())
	}
	fmt.Fprintf(w, "Content in file is...\r\n%s", string(dat))
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/read", readhandler)
	http.ListenAndServe(":8080", nil)
}
