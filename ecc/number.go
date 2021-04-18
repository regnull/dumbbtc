package ecc

import (
	"math"
	"math/big"
)

type Number interface {
	IsZero() bool
	Equal(Number) bool
	NotEqual(Number) bool
	Add(Number) Number
	Sub(Number) Number
	Mul(Number) Number
	MulScalar(*big.Int) Number
	Pow(*big.Int) Number
	Div(Number) Number
}

type Int64 struct {
	num int64
}

func NewInt64(num int64) *Int64 {
	return &Int64{num: num}
}

func (i *Int64) IsZero() bool {
	return i.num == 0
}

func (i *Int64) Equal(other Number) bool {
	otherInt64, ok := other.(*Int64)
	if !ok {
		return false
	}
	return i.num == otherInt64.num
}

func (i *Int64) NotEqual(other Number) bool {
	return !i.Equal(other)
}

func (i *Int64) Add(other Number) Number {
	otherInt64, ok := other.(*Int64)
	if !ok {
		panic("wrong type")
	}
	return &Int64{num: i.num + otherInt64.num}
}

func (i *Int64) Sub(other Number) Number {
	otherInt64, ok := other.(*Int64)
	if !ok {
		panic("wrong type")
	}
	return &Int64{num: i.num - otherInt64.num}
}

func (i *Int64) Mul(other Number) Number {
	otherInt64, ok := other.(*Int64)
	if !ok {
		panic("wrong type")
	}
	return &Int64{num: i.num * otherInt64.num}
}

func (i *Int64) MulScalar(x *big.Int) Number {
	return &Int64{num: i.num * x.Int64()}
}

func (i *Int64) Pow(exp *big.Int) Number {
	return &Int64{num: int64(math.Pow(float64(i.num), float64(exp.Int64())))}
}

func (i *Int64) Div(other Number) Number {
	otherInt64, ok := other.(*Int64)
	if !ok {
		panic("wrong type")
	}
	return &Int64{num: i.num / otherInt64.num}
}
