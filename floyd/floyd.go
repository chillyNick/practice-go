package floyd

// Triangle makes a Floyd's triangle matrix with rows count.
func Triangle(rows int) [][]int {
	res := make([][]int, rows)
	c := 0
	for i := 0; i < rows; i++ {
		res[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			c++
			res[i][j] = c
		}
	}

	return res
}
