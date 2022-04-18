package coins

import "math"

func Piles(n int) int {
	return calc(n, n)
}

func calc(maxSize, avlCoins int) int {
	if avlCoins == 0 {
		return 1
	}
	r := 0
	for c := int(math.Min(float64(avlCoins), float64(maxSize))); c > 0; c-- {
		r += calc(c, avlCoins-c)
	}

	return r
}
