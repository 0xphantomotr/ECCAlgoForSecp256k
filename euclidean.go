package main

import (
	"math/big"
)

func ExtendedGCD(a, b *big.Int) (*big.Int, *big.Int, *big.Int) {
	u := []*big.Int{big.NewInt(1), big.NewInt(0), new(big.Int).Set(a)}
	v := []*big.Int{big.NewInt(0), big.NewInt(1), new(big.Int).Set(b)}
	q := new(big.Int)
	temp := new(big.Int)

	for v[2].Sign() != 0 {
		q.Quo(u[2], v[2])

		u, v = v, u

		for i := range v {
			temp.Mul(q, v[i])
			u[i].Sub(u[i], temp)
		}
	}

	if u[1].Sign() < 0 {
		u[1].Add(u[1], b)
	}

	return u[2], u[0], u[1]
}

func ModInverse(a, mod *big.Int) *big.Int {
	_, x, _ := ExtendedGCD(a, mod)
	return x.Mod(x, mod)
}
