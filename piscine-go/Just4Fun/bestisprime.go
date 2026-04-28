package piscine

import "math/big"

func BestIsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	if nb <= 3 {
		return true
	}
	if nb%2 == 0 {
		return false
	}

	// Quick trial division by small primes to catch easy composites
	smallPrimes := []int{3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37}
	for _, p := range smallPrimes {
		if nb == p {
			return true
		}
		if nb%p == 0 {
			return false
		}
	}

	// Deterministic Miller–Rabin bases for testing 64-bit integers
	bases := []int64{2, 325, 9375, 28178, 450775, 9780504, 1795265022}

	n := big.NewInt(int64(nb))
	if n.Cmp(big.NewInt(2)) == 0 {
		return true
	}

	// write n-1 as d*2^s with d odd
	d := new(big.Int).Sub(n, big.NewInt(1))
	s := 0
	for d.Bit(0) == 0 {
		d.Rsh(d, 1)
		s++
	}

	one := big.NewInt(1)
	nMinusOne := new(big.Int).Sub(n, one)

	for _, a := range bases {
		aBig := big.NewInt(a)
		if aBig.Cmp(n) >= 0 {
			continue
		}
		x := new(big.Int).Exp(aBig, d, n)
		if x.Cmp(one) == 0 || x.Cmp(nMinusOne) == 0 {
			continue
		}
		composite := true
		for r := 1; r < s; r++ {
			x.Exp(x, big.NewInt(2), n)
			if x.Cmp(nMinusOne) == 0 {
				composite = false
				break
			}
		}
		if composite {
			return false
		}
	}
	return true
}
