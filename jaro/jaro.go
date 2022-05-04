package jaro

import (
	"strings"
)

func Distance(word1 string, word2 string) float64 {
	word1 = strings.ToLower(word1)
	word2 = strings.ToLower(word2)

	if word1 == word2 {
		return 1
	}

	l1, l2 := len(word1), len(word2)
	max := l1
	if l2 > max {
		max = l2
	}

	m1, m2 := make([]bool, l1), make([]bool, l2)
	md := max/2 - 1
	m := 0
	for i := 0; i < l1; i++ {
		sj := i - md
		if sj < 0 {
			sj = 0
		}

		for j := sj; j < l2 && j <= i+md; j++ {
			if !m2[j] && word1[i] == word2[j] {
				m1[i] = true
				m2[j] = true
				m++

				break
			}
		}
	}

	if m == 0 {
		return 0
	}

	t := 0
	j := 0
	for i, v := range m1 {
		if !v {
			continue
		}
		for !m2[j] {
			j++
		}

		if word1[i] != word2[j] {
			t++
		}

		j++
	}
	t /= 2

	mf := float64(m)

	return (mf/float64(l1) + mf/float64(l2) + (mf-float64(t))/mf) / 3
}
