package main

import (
	"math/big"
)

func TonelliShanks(n, p *big.Int) *big.Int {
	if !EulerCriterion(n, p) {
		return nil
	}

	Q := new(big.Int).Sub(p, big.NewInt(1))
	S := uint(0)
	for Q.Bit(0) == 0 {
		S++
		Q.Rsh(Q, 1)
	}

	z := big.NewInt(2)
	for EulerCriterion(z, p) {
		z.Add(z, big.NewInt(1))
	}

	M := S
	c := ModPow(z, Q, p)
	t := ModPow(n, Q, p)
	R := ModPow(n, new(big.Int).Add(Q, big.NewInt(1)).Rsh(Q, 1), p)

	for t.Cmp(big.NewInt(1)) != 0 {
		if t.Sign() == 0 {
			return big.NewInt(0)
		}

		t2 := new(big.Int).Exp(t, big.NewInt(2), p)
		i := uint(1)
		for t2.Cmp(big.NewInt(1)) != 0 && i < M {
			t2.Exp(t2, big.NewInt(2), p)
			i++
		}

		if i == M {
			return nil
		}

		b := ModPow(c, new(big.Int).Lsh(big.NewInt(1), M-i-1), p)
		c.Mul(b, b).Mod(c, p)
		t.Mul(t, c).Mod(t, p)
		R.Mul(R, b).Mod(R, p)
		M = i
	}

	return R
}
