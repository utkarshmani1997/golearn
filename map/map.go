package main

import (
	"fmt"
)

type myMap map[string]string
type myStruct struct {
	a int
}

func (m myMap) contains(key string) (val string, ok bool) {
	val, ok = m[key]
	return val, ok
}

func (m myMap) add(key, value string) {
	m[key] = value
}

func main() {
	// non idiomatic
	//var m = make(map[string]string)
	var m = map[string]string{}
	// both of the above way of initializing map is correct
	// make gives you more flexibility of preallocating memory
	// when you know that map will not grow after some time.So
	// there is little performance benifit
	m["x"] = "hi"
	//	m = nil
	delete(m, "y")
	if v, ok := m["x"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("not found")
	}

	m["jai"] = "bhai"
	m["jaii"] = "bhai"

	fmt.Println(m)

	var mymap = map[string]myStruct{}
	mymap["hi"] = myStruct{1}
	//	mymap["hi"].a = 2 cannot assign to struct field error (invalid)

	// idiomatic
	var idiomaticM = make(myMap)
	idiomaticM.add("x", "hi")
	if v, ok := idiomaticM.contains("x"); ok {
		fmt.Println(v)
	} else {
		fmt.Println("not found")
	}
}
