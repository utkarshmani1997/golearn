package main

import "fmt"

type myMap map[string]string

func (m myMap) contains(key string) (val string, ok bool) {
	val, ok = m[key]
	return val, ok
}

func (m myMap) add(key, value string) {
	m[key] = value
}

func main() {
	// non idiomatic
	var m = make(map[string]string)
	m["x"] = "hi"
	m = nil
	delete(m, "y")
	if v, ok := m["x"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("not found")
	}

	// idiomatic
	var idiomaticM = make(myMap)
	idiomaticM.add("x", "hi")
	if v, ok := idiomaticM.contains("x"); ok {
		fmt.Println(v)
	} else {
		fmt.Println("not found")
	}
}
