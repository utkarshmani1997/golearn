package biscuit

type Biscuit interface {
	Mix() Biscuit
	Make() Biscuit
	Pack() string
}
