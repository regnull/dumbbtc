package ecc

import (
	"fmt"
	"math/big"
)

type FieldElement struct {
	num   *big.Int
	prime *big.Int
}

var zero = big.NewInt(0)

func NewFieldElement(num, prime *big.Int) *FieldElement {
	if num.Cmp(prime) >= 0 || num.Cmp(zero) < 0 {
		panic("bad number")
	}
	return &FieldElement{num: num, prime: prime}
}

func NewFieldElementFromInt(num, prime int64) *FieldElement {
	return NewFieldElement(big.NewInt(num), big.NewInt(prime))
}

func (fe *FieldElement) Num() *big.Int {
	return fe.num
}

func (fe *FieldElement) String() string {
	return fmt.Sprintf("FieldElement_%s (%s)", fe.num.String(), fe.prime.String())
}

func (fe *FieldElement) IsZero() bool {
	return fe.num.Cmp(zero) == 0
}

func (fe *FieldElement) Equal(other Number) bool {
	otherFieldElement, ok := other.(*FieldElement)
	if !ok {
		return false
	}

	return fe.num.Cmp(otherFieldElement.num) == 0 && fe.prime.Cmp(otherFieldElement.prime) == 0
}

func (fe *FieldElement) NotEqual(other Number) bool {
	return !fe.Equal(other)
}

func (fe *FieldElement) Add(other Number) Number {
	otherFieldElement, ok := other.(*FieldElement)
	if !ok {
		panic("wrong type")
	}

	if fe.prime.Cmp(otherFieldElement.prime) != 0 {
		panic("bad number")
	}

	r := new(big.Int)
	r.Add(fe.num, otherFieldElement.num)
	r.Mod(r, fe.prime)
	return NewFieldElement(r, fe.prime)
}

func (fe *FieldElement) Sub(other Number) Number {
	otherFieldElement, ok := other.(*FieldElement)
	if !ok {
		panic("wrong type")
	}

	r := new(big.Int)
	r.Sub(fe.num, otherFieldElement.num)
	r.Mod(r, fe.prime)
	return NewFieldElement(r, fe.prime)
}

func (fe *FieldElement) Mul(other Number) Number {
	otherFieldElement, ok := other.(*FieldElement)
	if !ok {
		panic("wrong type")
	}

	if fe.prime.Cmp(otherFieldElement.prime) != 0 {
		panic("bad number")
	}
	r := new(big.Int)
	r.Mul(fe.num, otherFieldElement.num)
	r.Mod(r, fe.prime)
	return NewFieldElement(r, fe.prime)
}

func (fe *FieldElement) MulScalar(x *big.Int) Number {
	r := new(big.Int)
	r.Mul(fe.num, x)
	r.Mod(r, fe.prime)
	return NewFieldElement(r, fe.prime)
}

func (fe *FieldElement) Pow(exp *big.Int) Number {
	n := new(big.Int)
	n.Mod(exp, new(big.Int).Sub(fe.prime, big.NewInt(1)))
	r := new(big.Int)
	r.Exp(fe.num, n, fe.prime)
	return NewFieldElement(r, fe.prime)
}

func (fe *FieldElement) Div(other Number) Number {
	otherFieldElement, ok := other.(*FieldElement)
	if !ok {
		panic("wrong type")
	}
	if fe.prime.Cmp(otherFieldElement.prime) != 0 {
		panic("bad number")
	}

	return fe.Mul(other.Pow(new(big.Int).Sub(otherFieldElement.prime, big.NewInt(2))))
}

func (fe *FieldElement) Copy() Number {
	n1 := new(big.Int)
	n1.Set(fe.num)
	p1 := new(big.Int)
	p1.Set(fe.prime)
	return NewFieldElement(n1, p1)
}
