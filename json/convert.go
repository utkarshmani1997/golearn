package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type ColorGroup1 struct {
		ID     float64 `json:"id,string"`
		Name   string
		Colors []string
	}
	type ColorGroup2 struct {
		ID     float64
		Name   string
		Colors []string
	}
	group1 := ColorGroup1{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.Marshal(group1)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
	fmt.Println()
	group2 := ColorGroup2{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	c, err := json.Marshal(group2)
	if err != nil {
		fmt.Println("error:", err)
	}

	os.Stdout.Write(c)
}
