package main

import "fmt"

type X struct {
	y []Z
}

type Y struct {
	y []*Z
}

type Z struct {
	x int
}

type W struct {
	w *int
	x []string
}

func (x *X) inc() {
	for i := range x.y {
		x.y[i].x++
	}
}

func (y *Y) inc() {
	for i := range y.y {
		y.y[i].x++
	}
}

func (x *X) assign(a int) {
	for i := range x.y {
		x.y[i].x = a
	}
}

func main() {
	x := X{
		y: []Z{
			{x: 1},
			{x: 2},
			{x: 3},
			{x: 4},
			{x: 5},
		},
	}

	z := []*Z{
		{x: 5},
		{x: 6},
		{x: 7},
		{x: 8},
		{x: 9},
	}

	y := Y{
		y: z,
	}

	//	x.assign(5)
	x.inc()
	y.inc()

	w := &W{}
	fmt.Println(x, "\n", y.y, w.w, len(w.x))
}
