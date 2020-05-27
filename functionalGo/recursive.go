package main

var f = func(i int) {
	print("x")
}

func main() {
	var f func(x int)
	f = func(i int) {
		print(i)
		if i > 0 {
			f(i - 1)
		}
	}
	f(10)
}
