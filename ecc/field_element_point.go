package ecc

type FieldElementPoint struct {
	a, b *FieldElement
	x, y *FieldElement
}

func NewFieldElementPoint(x, y *FieldElement, a, b *FieldElement) *FieldElementPoint {
	if x == nil || y == nil || a == nil || b == nil {
		panic("bad arguments")
	}
	if y.Pow(2).NotEqual(
		x.Pow(3).Add(
			x.Multiply(a).Add(b))) {
		panic("point not on curve")
	}
	return &FieldElementPoint{x: x, y: y, a: a, b: b}
}

func NewFieldElementInf(a, b *FieldElement) *FieldElementPoint {
	if a == nil || b == nil {
		panic("bad arguments")
	}
	return &FieldElementPoint{x: nil, y: nil, a: a, b: b}
}

func (p *FieldElementPoint) IsInf() bool {
	return p.x == nil && p.y == nil
}

func (p *FieldElementPoint) Equal(other *FieldElementPoint) bool {
	return p.x.Equal(other.x) && p.y.Equal(other.y) && p.a.Equal(other.a) && p.b.Equal(other.b)
}

func (p *FieldElementPoint) NotEqual(other *FieldElementPoint) bool {
	return !p.Equal(other)
}

func (p *FieldElementPoint) Add(other *FieldElementPoint) *FieldElementPoint {
	if p.a.NotEqual(other.a) || p.b.NotEqual(other.b) {
		panic("points are not on the same curve")
	}
	if p.IsInf() {
		return &FieldElementPoint{a: other.a, b: other.b, x: other.x, y: other.y}
	}
	if other.IsInf() {
		return &FieldElementPoint{a: p.a, b: p.b, x: p.x, y: p.y}
	}

	if p.x.Equal(other.x) {
		return NewFieldElementInf(p.a, p.b)
	}

	if p.Equal(other) {
		// Two points are the same.

		if p.y.Equal(NewFieldElement(0, p.y.prime)) {
			// Vertical line.
			return NewFieldElementInf(p.a, p.b)
		}

		s := NewFieldElement(3, p.x.prime).Multiply(p.x).Multiply(p.x).Divide(NewFieldElement(2, p.x.prime)).Divide(p.y)
		x := s.Multiply(s).Subtract(NewFieldElement(2, p.x.prime).Multiply(p.x))
		y := s.Multiply(p.x.Subtract(x)).Subtract(p.y)

		return NewFieldElementPoint(x, y, p.a, p.b)
	}

	s := other.y.Subtract(p.y).Divide(other.x.Subtract(p.x))
	x := s.Multiply(s).Subtract(other.x).Subtract(p.x)
	y := s.Multiply(p.x.Subtract(x)).Subtract(p.y)

	return NewFieldElementPoint(x, y, p.a, p.b)
}
