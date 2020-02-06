package graph_basic

import "data"

func TopoSort(g [][]data.Edge) []int {
	ret := make([]int, len(g))
	c := len(g) - 1
	DFS(g, func(i int) {
		ret[c] = i
		c--
	})
	return ret
}
