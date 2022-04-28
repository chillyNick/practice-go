package buildword

import "sort"

type frsBasket struct {
	i   int
	frL []int
}

func BuildWord(word string, fragments []string) int {
	sort.Slice(fragments, func(i, j int) bool {
		return len(fragments[i]) > len(fragments[j])
	})

	p := make([]frsBasket, len(word))
	for _, f := range fragments {
		for _, i := range findIndexes(f, word) {
			if len(p[i].frL) == 0 {
				p[i].frL = []int{len(f)}
			} else {
				p[i].frL = append(p[i].frL, len(f))
			}
		}
	}

	c := 0
	lfl := 0
	for i := 0; i != len(word); {
		if i < 0 {
			c = 0
			break
		}

		if i > len(word) || p[i].i >= len(p[i].frL) {
			i -= lfl
			c--
			continue
		}

		lfl = p[i].frL[p[i].i]
		p[i].i += 1
		i += lfl
		c++
	}

	return c
}

func findIndexes(needle, haystack string) []int {
	res := make([]int, 0)
	for i := 0; i <= len(haystack)-len(needle); i++ {
		contain := true
		for j := 0; j < len(needle); j++ {
			if needle[j] != haystack[i+j] {
				contain = false
				break
			}
		}
		if contain {
			res = append(res, i)
		}
	}

	return res
}
