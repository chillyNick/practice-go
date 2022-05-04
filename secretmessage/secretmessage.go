package secretmessage

type item struct {
	v rune
	c int
	i int
}

// Decode func
func Decode(encoded string) string {
	m := make(map[rune]*item)
	s := make([]*item, 0, 256)
	for _, v := range encoded {
		itm, seen := m[v]
		if !seen {
			m[v] = &item{v, 1, len(s)}
			s = append(s, m[v])
			continue
		}

		itm.c++
		i := itm.i - 1
		for i >= 0 && s[i].c < itm.c {
			i--
		}
		i++

		s[i].i, itm.i = itm.i, s[i].i
		s[itm.i], s[s[i].i] = s[s[i].i], s[itm.i]
	}

	d := []rune{}
	for _, i := range s {
		if i.v == '_' {
			break
		}

		d = append(d, i.v)
	}

	return string(d)
}
