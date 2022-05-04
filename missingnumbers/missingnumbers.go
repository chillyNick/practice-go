package missingnumbers

func Missing(numbers []int) []int {
	fls := make([]bool, len(numbers)+3)

	for _, v := range numbers {
		fls[v] = true
	}

	res := make([]int, 0, 2)
	for i := 1; i < len(fls); i++ {
		if !fls[i] {
			res = append(res, i)
		}
	}

	return res
}
