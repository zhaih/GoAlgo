package test

import (
	"algo/dijkstra"
	"data"
	"testing"
)

func TestDijkstra(t *testing.T) {
	input := []data.Edge{
		{0, 1, 1},
		{0, 2, 6},
		{0, 4, 5},
		{1, 4, 2},
		{1, 3, 8},
		{2, 5, 1},
		{4, 2, 2},
		{4, 3, 3},
		{4, 5, 7},
		{5, 3, 1},
	}
	ans := []int{0, 1, 5, 6, 3, 6}
	g := data.Edge2Adj(input, 6, true)
	ret := dijkstra.Dijkstra(g, 0)
	for i, v := range ans {
		if v != ret[i] {
			t.Errorf("Not Equal. %d %d at %d", v, ret[i], i)
		}
	}
}
