package main

import (
	"fmt"
)

type Vehicle interface {
	NoOfGears() int
	NoOfTyres() int
	passengers() int
	typ() string
}

type specs struct {
	gears, tyres int
	ty           string
}
type Car struct {
	maxNoOfPassengers int
	spec              specs
}

type Bike struct {
	maxPassengers int
	spec          specs
}

type Airplane struct {
	passenger int
	spec      specs
}

func (c Car) NoOfGears() int {
	return c.spec.gears
}

func (c Car) NoOfTyres() int {
	return c.spec.tyres
}

func (c Car) passengers() int {
	return c.maxNoOfPassengers
}

func (c Car) typ() string {
	return c.spec.ty
}

func (b Bike) NoOfGears() int {
	return b.spec.gears
}

func (b Bike) NoOfTyres() int {
	return b.spec.tyres
}

func (b Bike) passengers() int {
	return b.maxPassengers
}

func (b Bike) typ() string {
	return b.spec.ty
}

func (a Airplane) NoOfGears() int {
	return a.spec.gears
}

func (a Airplane) NoOfTyres() int {
	return a.spec.tyres
}

func (a Airplane) passengers() int {
	return a.passenger
}

func (a Airplane) typ() string {
	return a.spec.ty
}

func displayInfo(v []Vehicle) {
	for _, pass := range v {
		fmt.Printf("No of passengers From %s = %d, which has %d no of tyres, %d no of gears \n", pass.typ(), pass.passengers(), pass.NoOfTyres(), pass.NoOfGears())
	}
}

func main() {
	vehicle1 := Car{spec: specs{gears: 6, tyres: 5, ty: "Car"}, maxNoOfPassengers: 5}
	vehicle2 := Bike{spec: specs{gears: 4, tyres: 2, ty: "Bike"}, maxPassengers: 2}
	vehicle3 := Airplane{spec: specs{gears: 25, tyres: 2, ty: "Airplane"}, passenger: 200}
	vehicles := []Vehicle{vehicle1, vehicle2, vehicle3}
	displayInfo(vehicles)
}
