package main //The package “main” tells the Go compiler that the package should compile as an executable program
//instead of a shared library.

import (
	"flag" //importing this package to do something with the command line arguments with the predefined methods
	"fmt"  //importing the package that contains some predefined methods.
	//existing in this package.
	"io/ioutil" //importing this package for reading and printing the source-code fetched.
	"net/http"  //importing this package for retrieving the source-code of the webpage requested.
	"os"        //importing this package for the system calls.
	"reflect"
)

func main() {

	//    var args string
	flag.Parse()
	args := flag.Args()
	//  fmt.Println("Enter the URL : ")
	//  fmt.Scanf("%s ",&args)
	fmt.Println(args)
	if len(args) < 1 {
		fmt.Println(reflect.TypeOf(args), "Please Enter the URL")
		os.Exit(1)
	}
	retrieve(args[0]) //call the retrieve function
}

func retrieve(url string) { //gives the  source code as output.

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("read error is:", err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read error is:", err)
		return
	} else {
		fmt.Println(string(body))

	}
}
