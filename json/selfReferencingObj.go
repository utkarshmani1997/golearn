package main

import (
	"encoding/json"
	"fmt"
)

type X struct {
	Y string `json:"y,omitempty"`
	W []X    `json:"w,omitempty"`
	Z string `json:"z,omitempty"`
}

func main() {
	var (
		obj X
		err error
		b   []byte
	)
	x := []byte(`
{
	"y": "hello",
	"w": [{
		"y": "world!",
		"w": [{
			"y": "welcome",
			"z": "see"
		}],
		"z": "magic"
	}],
	"z": "here"
}
	`)
	err = json.Unmarshal(x, &obj)
	if err != nil {
		panic(err)
	}
	//	fmt.Printf("%+v\n", obj)

	b, err = json.MarshalIndent(obj, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b)
}
