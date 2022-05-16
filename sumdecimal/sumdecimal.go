package sumdecimal

import (
	"math"
	"math/big"
)

func SumDecimal(c int) int {
	if c < 0 {
		return 0
	}

	sum := int64(0)

	zero := big.NewInt(0)
	one := big.NewInt(1)
	ten := big.NewInt(10)
	hundred := big.NewInt(100)

	x := big.NewInt(0)
	t := big.NewInt(0)
	p := big.NewInt(int64(math.Sqrt(float64(c))))
	r := big.NewInt(int64(c))
	r.Sub(r, big.NewInt(0).Mul(p, p))
	p.Add(p, p)
	for i := 0; i < 1000 && r.Cmp(zero) != 0; i++ {
		r.Mul(r, hundred)
		p.Mul(p, ten)

		x.Div(r, p)
		t.Mul(t.Add(p, x), x)
		if t.Cmp(r) == 1 {
			x.Sub(x, one)
			t.Mul(t.Add(p, x), x)
		}
		r.Sub(r, t)

		sum += x.Int64()
		p.Add(p, big.NewInt(0).Add(x, x))
	}

	return int(sum)
}
