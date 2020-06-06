package main

import "fmt"

// remove special character from string
func main() {
	var x = "Wel#come$ T&o S@quad"
	var xx = []rune(x)
	/*
		for index, r := range xx {
			if r >= 97 && r <= 122 {
			} else if r >= 65 && r <= 90 {
			} else if r == 32 {
			} else {
				xx = append(xx[:index], xx[index+1:]...)
			}
		}
	*/

	for index, r := range xx {
		if !(r >= 97 && r <= 122) && !(r >= 65 && r <= 90) && !(r == 32) {
			xx = append(xx[:index], xx[index+1:]...)
		}
	}

	fmt.Println(string(xx))
}
