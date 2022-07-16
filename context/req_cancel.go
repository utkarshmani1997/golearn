package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
	"io"
    "io/ioutil"
)

func handler2(w http.ResponseWriter, r *http.Request) {
    go func(done <-chan struct{}) {
        <-done
        fmt.Println("message", "client connection has gone away, request got cancelled")
    }(r.Context().Done())

    io.Copy(ioutil.Discard, r.Body) // <-- read the body
    time.Sleep(30 * time.Second)
    fmt.Fprintf(w, "Hi there, I love %s!\n", r.URL.Path[1:])
}

func handler(w http.ResponseWriter, r *http.Request) {
    go func(done <-chan struct{}) {
        <-done
        fmt.Println("message", "client connection has gone away, request got cancelled")
    }(r.Context().Done())

    time.Sleep(10 * time.Second)
    fmt.Fprintf(w, "Hi there, I love %s!\n", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/hello", handler2)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
