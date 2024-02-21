package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type EllipticCurve struct {
	A, B, P *big.Int
}

type Point struct {
	X, Y  *big.Int
	Inf   bool
	Curve *EllipticCurve
}

func NewEllipticCurve(a, b, p *big.Int) *EllipticCurve {
	curve := &EllipticCurve{A: a, B: b, P: p}
	if !curve.checkCurveParams() {
		panic("invalid curve parameters")
	}
	return curve
}

func (curve *EllipticCurve) checkCurveParams() bool {
	zero := big.NewInt(0)
	if curve.A.Cmp(zero) < 0 || curve.A.Cmp(curve.P) >= 0 {
		return false
	}
	if curve.B.Cmp(zero) < 0 || curve.B.Cmp(curve.P) >= 0 {
		return false
	}
	if curve.P.Cmp(big.NewInt(1)) <= 0 {
		return false
	}
	return true
}

func (curve *EllipticCurve) GenRandPoint() *Point {
	for {
		x, err := rand.Int(rand.Reader, curve.P)
		if err != nil {
			continue
		}

		f := new(big.Int).Exp(x, big.NewInt(3), curve.P)
		f.Add(f, new(big.Int).Mul(curve.A, x))
		f.Add(f, curve.B)
		f.Mod(f, curve.P)

		if EulerCriterion(f, curve.P) {
			y := TonelliShanks(f, curve.P)
			return NewPoint(x, y, curve)
		} else {
			fmt.Println("f(x) is not a quadratic residue, retrying...")
		}
	}
}

func NewPoint(x, y *big.Int, curve *EllipticCurve) *Point {
	return &Point{
		X:     x,
		Y:     y,
		Inf:   false,
		Curve: curve,
	}
}
