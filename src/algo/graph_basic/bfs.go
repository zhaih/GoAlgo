package graph_basic

import (
	"algo/union_find"
	"data"
)

func BFS(adj [][]data.Edge, f func(int)) {
	uf := union_find.UnionFind{}
	uf.Init(len(adj))
	queue := make([]int, 0)
	explore := func(i int) {
		uf.Union(0, i)
		for _, e := range adj[i] {
			if !uf.IsUnion(i, e.To) {
				queue = append(queue, e.To)
			}
		}
		f(i)
	}
	explore(0)
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		explore(node)
	}
	for i := range adj {
		if !uf.IsUnion(0, i) {
			explore(i)
			for len(queue) != 0 {
				node := queue[0]
				queue = queue[1:]
				explore(node)
			}
		}
	}
}
