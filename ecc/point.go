package ecc

import (
	"fmt"
	"math/big"
)

type Point struct {
	a Number
	b Number
	x Number
	y Number
}

func NewPoint(x, y, a, b Number) *Point {
	if y.Pow(big.NewInt(2)).NotEqual(x.Pow(big.NewInt(3)).Add(a.Mul(x)).Add(b)) {
		panic("point not on curve")
	}

	return &Point{x: x, y: y, a: a, b: b}
}

func NewPointCopy(p *Point) *Point {
	a1 := p.a.Copy()
	b1 := p.b.Copy()
	x1 := p.x.Copy()
	y1 := p.y.Copy()
	return NewPoint(x1, y1, a1, b1)
}

func NewInf(a, b Number) *Point {
	return &Point{x: nil, y: nil, a: a, b: b}
}

func (p *Point) String() string {
	if p.IsInf() {
		return fmt.Sprintf("Point: inf")
	}
	return fmt.Sprintf("Point: %s, %s", p.x.String(), p.y.String())
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

		return NewPoint(x, y, p.a, p.b)
	}

	if p.x.Equal(other.x) {
		// Two points form a vertical line.
		return NewInf(p.a, p.b)
	}

	s := other.y.Sub(p.y).Div(other.x.Sub(p.x))
	x := s.Pow(big.NewInt(2)).Sub(other.x).Sub(p.x)
	y := s.Mul(p.x.Sub(x)).Sub(p.y)
	return NewPoint(x, y, p.a, p.b)
}

func (p *Point) Mul(factor *big.Int) *Point {
	f := new(big.Int)
	f.Set(factor)
	current := NewPointCopy(p)
	res := NewInf(p.a, p.b)
	zero := big.NewInt(0)
	one := big.NewInt(1)
	for f.Cmp(zero) > 0 {
		x := new(big.Int)
		if x.And(f, one).Cmp(zero) != 0 {
			res = res.Add(current)
		}
		current = current.Add(current)
		f.Rsh(f, 1)
	}
	return res
}
