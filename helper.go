package main

import (
	"math/big"
)

func ReverseBits(m *big.Int) (*big.Int, uint) {
	mRev := new(big.Int)
	n := uint(0)
	for mBit := new(big.Int).Set(m); mBit.Sign() > 0; n++ {
		mRev.Lsh(mRev, 1)
		mRev.Or(mRev, new(big.Int).And(mBit, big.NewInt(1)))
		mBit.Rsh(mBit, 1)
	}
	return mRev, n
}

func ModPow(base, exponent, modulus *big.Int) *big.Int {
	exponentReversed, bitLen := ReverseBits(exponent)

	result := big.NewInt(1)
	for i := uint(0); i < bitLen; i++ {
		if exponentReversed.Bit(int(i)) == 1 {
			result.Mul(result, base).Mod(result, modulus)
		}
		if i < bitLen-1 {
			result.Mul(result, result).Mod(result, modulus)
		}
	}

	return result
}

func OnSameCurve(points ...*Point) bool {
	for i := 0; i < len(points)-1; i++ {
		if points[i].Curve != points[i+1].Curve {
			return false
		}
	}
	return true
}

func (curve *EllipticCurve) IsOnCurve(point *Point) bool {
	if point.Inf {
		return true
	}
	ySquared := new(big.Int).Mul(point.Y, point.Y)
	ySquared.Mod(ySquared, curve.P)

	xCubed := new(big.Int).Exp(point.X, big.NewInt(3), curve.P)
	ax := new(big.Int).Mul(curve.A, point.X)
	ax.Mod(ax, curve.P)

	rhs := new(big.Int).Add(xCubed, ax)
	rhs.Add(rhs, curve.B)
	rhs.Mod(rhs, curve.P)

	return ySquared.Cmp(rhs) == 0
}

func EulerCriterion(n, p *big.Int) bool {
	exp := new(big.Int).Sub(p, big.NewInt(1))
	exp.Div(exp, big.NewInt(2))
	return ModPow(n, exp, p).Cmp(big.NewInt(1)) == 0
}
