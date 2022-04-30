package mergesort

// MergeSort is used to sort an array of integer
func MergeSort(input []int) []int {
	merge(input, make([]int, len(input)))

	return input
}

func merge(input []int, buff []int) {
	if len(input) < 2 {
		return
	}

	m := len(input) / 2
	merge(input[0:m], buff)
	merge(input[m:], buff)
	k, i, j := 0, 0, m
	for ; i < m && j < len(input); k++ {
		if input[i] < input[j] {
			buff[k] = input[i]
			i++
		} else {
			buff[k] = input[j]
			j++
		}
	}

	for ; i < m; i, k = i+1, k+1 {
		buff[k] = input[i]
	}

	for ; j < len(input); j, k = j+1, k+1 {
		buff[k] = input[j]
	}

	copy(input, buff)
}
