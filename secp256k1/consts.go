package secp256k1

import "math/big"

type CurveParamsStruct struct {
	P  *big.Int
	A  *big.Int
	B  *big.Int
	Gx *big.Int
	Gy *big.Int
	N  *big.Int
}

var CurveParams *CurveParamsStruct

func init() {
	CurveParams = &CurveParamsStruct{}
	CurveParams.P, _ = new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F", 16)
	CurveParams.A = big.NewInt(0)
	CurveParams.B = big.NewInt(7)
	CurveParams.Gx, _ = new(big.Int).SetString("79BE667EF9DCBBAC55A06295CE870B07029BFCDB2DCE28D959F2815B16F81798", 16)
	CurveParams.Gy, _ = new(big.Int).SetString("483ADA7726A3C4655DA4FBFC0E1108A8FD17B448A68554199C47D08FFB10D4B8", 16)
	CurveParams.N, _ = new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141", 16)
}
