package main

import (
	"fmt"
	"reflect"

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

	fooType := reflect.TypeOf(sweet.Sweet{})
	fmt.Println(fooType.Method(0))
	for i := 0; i < fooType.NumMethod(); i++ {
		method := fooType.Method(i)
		fmt.Println(method.Name)
	}

}
