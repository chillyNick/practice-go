package lastlettergame

import (
	"unicode/utf8"
)

type node struct {
	val    string
	childs []*node
	fr     rune
	lr     rune
}

func Sequence(dic []string) []string {
	nodes := make([]*node, 0, len(dic))
	for _, w := range dic {
		n := &node{
			val:    w,
			childs: make([]*node, 0),
		}
		n.fr, _ = utf8.DecodeRuneInString(n.val)
		n.lr, _ = utf8.DecodeLastRuneInString(n.val)

		for _, an := range nodes {
			if an.lr == n.fr {
				an.childs = append(an.childs, n)
			}
			if n.lr == an.fr {
				n.childs = append(n.childs, an)
			}
		}

		nodes = append(nodes, n)
	}

	return getLongestSeq(nodes)
}

func getLongestSeq(nodes []*node) []string {
	visited := make(map[*node]bool)
	seq := make([]string, 0)
	for _, v := range dfs(&node{childs: nodes}, &visited)[1:] {
		seq = append(seq, v.val)
	}

	return seq
}

func dfs(n *node, visited *map[*node]bool) []*node {
	(*visited)[n] = true
	var maxTail []*node
	for _, chN := range n.childs {
		if (*visited)[chN] {
			continue
		}

		s := dfs(chN, visited)
		if len(s) > len(maxTail) {
			maxTail = s
		}
	}

	(*visited)[n] = false

	return append([]*node{n}, maxTail...)
}
