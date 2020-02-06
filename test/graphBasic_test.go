package test

import (
	"algo/graph_basic"
	"data"
	"testhelper"
	"testing"
)

func TestBFS(t *testing.T) {
	es := data.GenBinaryTree(7)
	g := data.Edge2Adj(es, 7, true)
	c := 0
	graph_basic.BFS(g, func(i int) {
		testhelper.AssertEqual(t, c, i)
		c++
	})
}

func TestDFS(t *testing.T) {
	es := data.GenBinaryTree(7)
	g := data.Edge2Adj(es, 7, true)
	c := 0
	ans := []int{3, 4, 1, 5, 6, 2, 0}
	graph_basic.DFS(g, func(i int) {
		testhelper.AssertEqual(t, ans[c], i)
		c++
	})
}

func TestTopoSort(t *testing.T) {
	es := []data.Edge{
		{0, 1, 1},
		{1, 2, 1},
		{1, 3, 1},
		{2, 3, 1},
		{3, 4, 1},
	}
	g := data.Edge2Adj(es, 5, true)
	ary := graph_basic.TopoSort(g)
	for i := 0; i < 5; i++ {
		testhelper.AssertEqual(t, ary[i], i)
	}
}
