package main

import (
	"math/big"
	"testing"
)

func TestGeneratingPoint(t *testing.T) {
	curve := NewEllipticCurve(big.NewInt(1), big.NewInt(1), big.NewInt(17))
	t.Log("Starting point generation...")
	P := curve.GenRandPoint()
	t.Logf("Generated point: (%v, %v)", P.X, P.Y)

	if P.Inf {
		t.Fatal("Generated point is at infinity.")
	}

	t.Log("Verifying if point is on the curve...")
	ySquared := new(big.Int).Mul(P.Y, P.Y).Mod(P.Y, curve.P)
	rightSide := new(big.Int).Exp(P.X, big.NewInt(3), nil)
	rightSide.Add(rightSide, new(big.Int).Mul(curve.A, P.X))
	rightSide.Add(rightSide, curve.B)
	rightSide.Mod(rightSide, curve.P)

	t.Logf("y^2: %v", ySquared)
	t.Logf("x^3 + Ax + B mod P: %v", rightSide)

	if ySquared.Cmp(rightSide) != 0 {
		t.Errorf("Generated point (%v, %v) is not on the curve y^2 = x^3 + %vx + %v mod %v", P.X, P.Y, curve.A, curve.B, curve.P)
	} else {
		t.Logf("Generated point (%v, %v) is on the curve y^2 = x^3 + %vx + %v mod %v", P.X, P.Y, curve.A, curve.B, curve.P)
	}
}
