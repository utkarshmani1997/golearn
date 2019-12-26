package salty

import (
	"fmt"

	"github.com/utkarshmani1997/golearn/patterns/builder/pkg/biscuit"
)

type BuildOpts func(*Salty)

type Salty struct {
	salt       int
	wheatFlour int
	oil        int
	bakingSoda int
}

func New(opts ...BuildOpts) *Salty {
	saltyBiscuit := new(Salty)
	for _, o := range opts {
		o(saltyBiscuit)
	}
	return saltyBiscuit
}

func WithSalt(quantity int) BuildOpts {
	return func(s *Salty) {
		s.salt = quantity
	}
}

func WithWheatFlour(quantity int) BuildOpts {
	return func(s *Salty) {
		s.wheatFlour = quantity
	}
}

func WithOil(quantity int) BuildOpts {
	return func(s *Salty) {
		s.oil = quantity
	}
}

func WithBakingSoda(quantity int) BuildOpts {
	return func(s *Salty) {
		s.bakingSoda = quantity
	}
}

func (s *Salty) Make() biscuit.Biscuit {
	fmt.Printf("We are using Salt: %vkg, Oil: %vkg, WheatFlour: %vkg, BakingSoda: %vkg\n", s.salt, s.oil, s.wheatFlour, s.bakingSoda)
	return s
}

func (s *Salty) Mix() biscuit.Biscuit {
	fmt.Printf("Total weight of biscuit: %v\n", s.salt+s.wheatFlour+s.bakingSoda+s.oil)
	return s
}

func (s *Salty) Pack() string {
	return "Salty biscuit"
}
