package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type X struct {
	I int
}

func unmarshalFile(file string, obj interface{}) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	if err = dec.Decode(obj); err != nil && err == io.EOF {
		fmt.Println(err)
		return nil
	}
	return err
}

func main() {

	var x = X{}
	fmt.Println(unmarshalFile("xyz", &x), x)
}
