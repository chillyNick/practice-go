package wordladder

type item struct {
	d int
	i int
}

type queue struct {
	buf  []item
	head int
	tail int
}

func (q *queue) init(size int) {
	q.buf = make([]item, size+1)
}

func (q *queue) isEmpty() bool {
	return q.tail == q.head
}

func (q *queue) push(d int, i int) {
	q.buf[q.tail] = item{d, i}

	q.tail = (q.tail + 1) % len(q.buf)
}

func (q *queue) pop() item {
	r := q.buf[q.head]

	q.head = (q.head + 1) % len(q.buf)

	return r
}

func WordLadder(from string, to string, dic []string) int {
	_, ok := find(to, dic)
	if !ok {
		return 0
	}
	_, ok = find(from, dic)
	if !ok {
		dic = append(dic, from)
	}

	edges := make([][]bool, len(dic))
	for i := 0; i < len(dic); i++ {
		edges[i] = make([]bool, len(dic))
	}

	for i, w1 := range dic {
		for j, w2 := range dic {
			if i == j || edges[i][j] {
				continue
			}
			if isEdgeExist(w1, w2) {
				edges[i][j] = true
				edges[j][i] = true
			}
		}
	}

	q := new(queue)
	q.init(len(dic))

	visited := make([]bool, len(dic))
	i, _ := find(from, dic)
	visited[i] = true
	q.push(1, i)
	for !q.isEmpty() {
		it := q.pop()
		if dic[it.i] == to {
			return it.d
		}

		for i, v := range edges[it.i] {
			if !v || visited[i] {
				continue
			}

			visited[it.i] = true
			q.push(it.d+1, i)
		}
	}

	return 0
}

func isEdgeExist(w1, w2 string) bool {
	d := 0
	for i := 0; i < len(w1); i++ {
		if w1[i] != w2[i] {
			d++
		}
	}

	return d == 1
}

func find(needle string, dic []string) (int, bool) {
	for i, w := range dic {
		if w == needle {
			return i, true
		}
	}

	return 0, false
}
