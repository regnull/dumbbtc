package ecc

import (
	"math/big"
)

type Point struct {
	a Number
	b Number
	x Number
	y Number
}

func NewPoint(x, y, a, b Number) *Point {
	if y.Pow(big.NewInt(2)).NotEqual(x.Pow(big.NewInt(3)).Add(a.Mul(x).Add(b))) {
		panic("point not on curve")
	}

	return &Point{x: x, y: y, a: a, b: b}
}

func NewInf(a, b Number) *Point {
	return &Point{x: nil, y: nil, a: a, b: b}
}

func (p *Point) IsInf() bool {
	return p.x == nil && p.y == nil
}

func (p *Point) Equal(other *Point) bool {
	return p.x.Equal(other.x) && p.y.Equal(other.y) && p.a.Equal(other.a) && p.b.Equal(other.b)
}

func (p *Point) NotEqual(other *Point) bool {
	return !p.Equal(other)
}

func (p *Point) Add(other *Point) *Point {
	if p.a.NotEqual(other.a) || p.b.NotEqual(other.b) {
		panic("points are not on the same curve")
	}
	if p.IsInf() {
		return &Point{a: other.a, b: other.b, x: other.x, y: other.y}
	}
	if other.IsInf() {
		return &Point{a: p.a, b: p.b, x: p.x, y: p.y}
	}

	if p.x.Equal(other.x) {
		// Two points form a vertical line.
		return NewInf(p.a, p.b)
	}

	if p.Equal(other) {
		// Two points are the same.

		if p.y.IsZero() {
			// Vertical line.
			return NewInf(p.a, p.b)
		}

		s := p.x.Pow(big.NewInt(2)).MulScalar(big.NewInt(3)).
			Add(p.a).Div(p.y.MulScalar(big.NewInt(2)))
		x := s.Pow(big.NewInt(2)).Sub(p.x.MulScalar(big.NewInt(2)))
		y := s.Mul(p.x.Sub(x)).Sub(p.y)

		//s := (3*p.x*p.x + p.a) / 2 / p.y
		//x := s*s - 2*p.x
		//y := s*(p.x-x) - p.y
		return NewPoint(x, y, p.a, p.b)
	}

	s := other.y.Sub(p.y).Div(other.x.Sub(p.x))
	//s := (other.y - p.y) / (other.x - p.x)
	x := s.Pow(big.NewInt(2)).Sub(other.x).Sub(p.x)
	//x := s*s - other.x - p.x
	y := s.Mul(p.x.Sub(x)).Sub(p.y)
	//y := s*(p.x-x) - p.y
	return NewPoint(x, y, p.a, p.b)
}
