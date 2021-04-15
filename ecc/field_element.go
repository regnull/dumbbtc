package ecc

import "fmt"

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
