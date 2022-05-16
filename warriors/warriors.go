package warriors

func Count(image string) int {
	field := make([][]int, 0)
	buf := make([]int, 0)
	for _, v := range image {
		switch v {
		case '1':
			buf = append(buf, 1)
		case '0':
			buf = append(buf, 0)
		case '\n':
			field = append(field, buf)
			buf = make([]int, 0)
		}
	}

	c := 2
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			if field[i][j] == 1 {
				mark(field, i, j, c)
				c++
			}
		}
	}

	return c - 2
}

func mark(field [][]int, i, j, m int) {
	if i < 0 || j < 0 || i >= len(field) || j >= len(field[0]) || field[i][j] != 1 {
		return
	}

	field[i][j] = m
	for ti := -1; ti < 2; ti++ {
		for tj := -1; tj < 2; tj++ {
			mark(field, ti+i, tj+j, m)
		}
	}
}
