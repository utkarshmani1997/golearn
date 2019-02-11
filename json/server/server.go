package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/utkarshmani1997/golearn/json/types"
)

func main() {
	http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
		var user types.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			fmt.Fprintf(w, "%v, %v", err, user)
		}
		err := json.NewEncoder(w).Encode(user)
		if err != nil {
			fmt.Fprintf(w, "%v, %v", err, user)
		}
		fmt.Println(user)
	})
	//	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":8080", nil)
	fmt.Println("err: ", err)
}
