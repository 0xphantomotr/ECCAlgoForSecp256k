package main

import (
	"math/big"
)

func (p *Point) Add(other *Point) *Point {
	if p.Inf {
		return other
	}
	if other.Inf {
		return p
	}
	if p.X.Cmp(other.X) == 0 && (p.Y.Cmp(other.Y) != 0 || p.Y.Sign() == 0) {
		return NewPoint(big.NewInt(0), big.NewInt(0), p.Curve).SetInfinity(true)
	}

	var lambda *big.Int
	if p.X.Cmp(other.X) == 0 {
		two := big.NewInt(2)
		three := big.NewInt(3)
		lambda = new(big.Int).Mul(three, new(big.Int).Exp(p.X, two, p.Curve.P))
		lambda.Add(lambda, p.Curve.A)
		temp := new(big.Int).Mul(two, p.Y)
		lambda.Mul(lambda, new(big.Int).ModInverse(temp, p.Curve.P))
	} else {
		lambda = new(big.Int).Sub(other.Y, p.Y)
		temp := new(big.Int).Sub(other.X, p.X)
		lambda.Mul(lambda, new(big.Int).ModInverse(temp, p.Curve.P))
	}

	lambda.Mod(lambda, p.Curve.P)

	xResult := new(big.Int).Exp(lambda, big.NewInt(2), p.Curve.P)
	xResult.Sub(xResult, p.X)
	xResult.Sub(xResult, other.X)
	xResult.Mod(xResult, p.Curve.P)

	yResult := new(big.Int).Sub(p.X, xResult)
	yResult.Mul(yResult, lambda)
	yResult.Sub(yResult, p.Y)
	yResult.Mod(yResult, p.Curve.P)

	return NewPoint(xResult, yResult, p.Curve)
}

func (p *Point) SetInfinity(inf bool) *Point {
	p.Inf = inf
	if inf {
		p.X, p.Y = nil, nil
	}
	return p
}

func (p *Point) Negate() *Point {
	if p.Inf {
		return p
	}
	negY := new(big.Int).Neg(p.Y)
	negY.Mod(negY, p.Curve.P)
	return NewPoint(p.X, negY, p.Curve)
}

func (p *Point) Subtract(other *Point) *Point {
	return p.Add(other.Negate())
}

func (p *Point) ScalarMult(k *big.Int) *Point {
	result := NewPoint(big.NewInt(0), big.NewInt(0), p.Curve).SetInfinity(true)
	tempPoint := NewPoint(new(big.Int).Set(p.X), new(big.Int).Set(p.Y), p.Curve)

	for kBit := k.BitLen() - 1; kBit >= 0; kBit-- {
		if k.Bit(kBit) == 1 {
			result = result.Add(tempPoint)
		}
		tempPoint = tempPoint.Add(tempPoint)
	}

	return result
}
