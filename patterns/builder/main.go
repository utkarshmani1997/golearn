package main

import (
	"fmt"

	"github.com/utkarshmani1997/golearn/patterns/builder/pkg/biscuit/salty"
	"github.com/utkarshmani1997/golearn/patterns/builder/pkg/biscuit/sweet"
)

func main() {
	sweetBiscuit := sweet.New().
		WithSugar(10).
		WithWheatFlour(15).
		WithOil(3).
		WithBakingSoda(1).
		Make().
		Mix().
		Pack()

	fmt.Println(sweetBiscuit)

	// Another variant of builder pattern using functional
	// approach
	saltyBiscuit := salty.New(
		salty.WithSalt(10),
		salty.WithWheatFlour(15),
		salty.WithOil(3),
		salty.WithBakingSoda(1)).
		Make().
		Mix().
		Pack()
	fmt.Println(saltyBiscuit)
}
