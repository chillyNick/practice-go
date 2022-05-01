package romannumerals

var m map[rune]int = map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

var ro []rune = []rune{'M', 'D', 'C', 'L', 'X', 'V', 'I'}

func Encode(n int) (string, bool) {
	if n <= 0 {
		return "", false
	}

	res := ""
	for i, cr := range ro {
		rv := m[cr]
		c := n / rv
		n %= rv

		if c > 4 {
			return "", false
		}
		if c == 4 {
			res += string(cr) + string(ro[i-1])
			continue
		}

		for c != 0 {
			res += string(cr)
			c--
		}
	}

	return res, true
}

func Decode(s string) (int, bool) {
	if s == "" {
		return 0, false
	}

	res := 0
	prv := 1000
	for _, v := range s {
		rv, ok := m[v]
		if !ok {
			return 0, false
		}

		if prv < rv {
			res += rv - prv - prv
		} else {
			res += rv
		}

		prv = rv
	}

	return res, true
}
