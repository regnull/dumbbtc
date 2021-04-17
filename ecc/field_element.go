package ecc

import (
	"fmt"
	"math"
)

type FieldElement struct {
	num   int
	prime int
}

func NewFieldElement(num, prime int) *FieldElement {
	if num >= prime || num < 0 {
		panic("bad number")
	}
	return &FieldElement{num: num, prime: prime}
}

func (fe *FieldElement) String() string {
	return fmt.Sprintf("FieldElement_%d (%d)", fe.num, fe.prime)
}

func (fe *FieldElement) Equal(x *FieldElement) bool {
	if x == nil {
		return false
	}
	return fe.num == x.num && fe.prime == x.prime
}

func (fe *FieldElement) NotEqual(x *FieldElement) bool {
	return !fe.Equal(x)
}

func (fe *FieldElement) Add(x *FieldElement) *FieldElement {
	if fe.prime != x.prime {
		panic("bad number")
	}
	return NewFieldElement(mod((fe.num+x.num), fe.prime), fe.prime)
}

func (fe *FieldElement) Subtract(x *FieldElement) *FieldElement {
	if fe.prime != x.prime {
		panic("bad number")
	}
	return NewFieldElement(mod((fe.num-x.num), fe.prime), fe.prime)
}

func (fe *FieldElement) Multiply(x *FieldElement) *FieldElement {
	if fe.prime != x.prime {
		panic("bad number")
	}
	return NewFieldElement(mod((fe.num*x.num), fe.prime), fe.prime)
}

func (fe *FieldElement) Pow(e int) *FieldElement {
	n := mod(e, (fe.prime - 1))
	return NewFieldElement(mod(int(math.Pow(float64(fe.num), float64(n))), fe.prime), fe.prime)
}

func (fe *FieldElement) Divide(x *FieldElement) *FieldElement {
	if fe.prime != x.prime || x.num == 0 {
		panic("bad number")
	}

	return fe.Multiply(x.Pow(x.prime - 2))
}
