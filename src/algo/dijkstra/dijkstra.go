package dijkstra

import (
	"algo/union_find"
	"data"
)

func Dijkstra(g [][]data.Edge, s int) []int {
	ret := make([]int, len(g))
	for i := 0; i < len(g); i++ {
		if i != s {
			ret[i] = 1 << 31
		}
	}
	uf := union_find.UnionFind{}
	uf.Init(len(g))
	next := s
	for next != -1 {
		u := next
		uf.Union(u, s)
		for _, e := range g[u] {
			// relaxing
			if ret[e.To] > ret[e.From]+e.W {
				ret[e.To] = ret[e.From] + e.W
			}
		}
		// should use fibonacci heap here, but not intend to implement that
		// use linear search instead
		next = -1
		min := 1 << 31
		for i := 0; i < len(g); i++ {
			if uf.IsUnion(i, s) {
				continue
			}
			if ret[i] < min {
				min = ret[i]
				next = i
			}
		}
	}
	return ret
}
