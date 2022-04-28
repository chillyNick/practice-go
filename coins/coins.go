package coins

import (
	"math"
	"strconv"
)

var buff map[string]int

func Piles(n int) int {
	if len(buff) == 0 {
		buff = make(map[string]int)
	}
	return calc(n, n)
}

func calc(maxSize, avlCoins int) int {
	if avlCoins == 0 {
		return 1
	}
	r := 0
	maxSize = int(math.Min(float64(avlCoins), float64(maxSize)))
	k := strconv.Itoa(maxSize) + ":" + strconv.Itoa(avlCoins)
	if rb, ok := buff[k]; ok {
		return rb
	}

	for c := maxSize; c > 0; c-- {
		r += calc(c, avlCoins-c)
	}

	buff[k] = r

	return r
}
