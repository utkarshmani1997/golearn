package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, the end point is :  %s !, query: %s", strings.TrimSpace(strings.TrimPrefix(r.URL.Path, "/hello/")), r.URL.RawQuery)
}

func main() {
	http.HandleFunc("/hello", handler)
	http.ListenAndServe(":8080", nil)
}
