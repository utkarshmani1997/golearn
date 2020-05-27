package main

import "fmt"

type X struct {
	a int
	b Y
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
	x.b = Y{
		b: 100,
	}
}

func main() {
	x := X{
		a: 5,
	}

	y := Y{
		b: 10,
	}

	//x.b = y

	//	x.xInc()
	//	y.yInc()
	Inc(&x)
	x.xInc()
	fmt.Println(x, y)
}
