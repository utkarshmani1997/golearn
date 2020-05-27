package main

import "fmt"

type myStruct struct {
	a int
	b map[string]hello
}

type hello struct {
	hi int
}

func main() {
	myy := myStruct{}
	m := map[string]myStruct{}
	fmt.Println(m["x"].a)
	fmt.Println(myy.b["x"].hi)
	my := myStruct{a: 2}
	//	my.b["hi"] = "hello"
	fmt.Println(my)
}
