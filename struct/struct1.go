package main

import "fmt"

type X struct {
	a int
}

type Y struct {
	b int
}

func (x *X) xInc() {
	x.a++
}

func (y Y) yInc() {
	y.b++
}

func Inc(x *X) {
	/*x = &X{
		a: 10,
	}
	*/
	x.a = 15
}

func main() {
	x := X{
		a: 5,
	}

	y := Y{
		b: 10,
	}

	//	x.xInc()
	//	y.yInc()
	Inc(&x)
	fmt.Println(x, y)
}
