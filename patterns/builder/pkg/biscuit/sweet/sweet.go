package sweet

import (
	"fmt"

	"github.com/utkarshmani1997/golearn/patterns/builder/pkg/biscuit"
)

type Sweet struct {
	sugar      int
	wheatFlour int
	oil        int
	bakingSoda int
}

func New() *Sweet {
	return new(Sweet)
}

func (s *Sweet) WithSugar(quantity int) *Sweet {
	s.sugar = quantity
	return s
}

func (s *Sweet) WithWheatFlour(quantity int) *Sweet {
	s.wheatFlour = quantity
	return s
}

func (s *Sweet) WithOil(quantity int) *Sweet {
	s.oil = quantity
	return s
}

func (s *Sweet) WithBakingSoda(quantity int) *Sweet {
	s.bakingSoda = quantity
	return s
}

func (s *Sweet) Make() biscuit.Biscuit {
	fmt.Printf("We are using Sugar: %vkg, Oil: %vkg, WheatFlour: %vkg, BakingSoda: %vkg\n", s.sugar, s.oil, s.wheatFlour, s.bakingSoda)
	return s
}

func (s *Sweet) Mix() biscuit.Biscuit {
	fmt.Printf("Total weight of biscuit: %v\n", s.sugar+s.wheatFlour+s.bakingSoda+s.oil)
	return s
}

func (s *Sweet) Pack() string {
	return "Sweet biscuit"
}
