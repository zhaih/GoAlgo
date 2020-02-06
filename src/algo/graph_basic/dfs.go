package graph_basic

import (
	"algo/union_find"
	"data"
)

func DFS(g [][]data.Edge, f func(int)) {
	uf := union_find.UnionFind{}
	uf.Init(len(g))
	var explore func(int)
	explore = func(i int) {
		uf.Union(0, i)
		for _, e := range g[i] {
			if !uf.IsUnion(e.From, e.To) {
				explore(e.To)
			}
		}
		f(i)
	}
	explore(0)
}
