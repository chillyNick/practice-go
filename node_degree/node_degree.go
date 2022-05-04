package nodedegree

import "fmt"

// Degree func
func Degree(nodes int, graph [][2]int, node int) (int, error) {
	if node > nodes {
		return 0, fmt.Errorf("node %v not found in the graph", node)
	}

	r := 0
	for i := range graph {
		if graph[i][0] == node || graph[i][1] == node {
			r++
		}
	}

	return r, nil
}
