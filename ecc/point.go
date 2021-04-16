package ecc

import "math"

type Point struct {
	a, b, x, y int
}

func NewPoint(x, y, a, b int) *Point {
	if y*y != x*x*x+a*x+b {
		panic("point not on curve")
	}
	return &Point{x: x, y: y, a: a, b: b}
}

func NewInf(a, b int) *Point {
	return &Point{x: math.MaxInt32, y: math.MaxInt32, a: a, b: b}
}

func (p *Point) IsInf() bool {
	return p.x == math.MaxInt32 && p.y == math.MaxInt32
}

func (p *Point) Equal(other *Point) bool {
	return p.x == other.x && p.y == other.y && p.a == other.a && p.b == other.b
}

func (p *Point) NotEqual(other *Point) bool {
	return !p.Equal(other)
}

func (p *Point) Add(other *Point) *Point {
	if p.a != other.a || p.b != other.b {
		panic("points are not on the same curve")
	}
	if p.IsInf() {
		return &Point{a: other.a, b: other.b, x: other.x, y: other.y}
	}
	if other.IsInf() {
		return &Point{a: p.a, b: p.b, x: p.x, y: p.y}
	}

	if p.x == other.x {
		return NewInf(p.a, p.b)
	}

	if p.Equal(other) {
		// Two points are the same.

		if p.y == 0 {
			// Vertical line.
			return NewInf(p.a, p.b)
		}

		s := (3*p.x*p.x + p.a) / 2 / p.y
		x := s*s - 2*p.x
		y := s*(p.x-x) - p.y
		return NewPoint(x, y, p.a, p.b)
	}

	s := (other.y - p.y) / (other.x - p.x)
	x := s*s - other.x - p.x
	y := s*(p.x-x) - p.y
	return NewPoint(x, y, p.a, p.b)
}
