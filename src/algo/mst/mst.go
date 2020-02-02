package mst

import (
	"algo/min_heap"
	"algo/union_find"
	"data"
)

type MSTInput struct {
	N    int     //max number of nodes
	Data [][]int //m x 3 2d array, where m is number of edges, each row contains (a,b,weight) (assume indirect edge)
}

/**
 @return: matrix representation
 */
func HandleMSTInput(in *MSTInput) [][]int {
	ret := make([][]int, in.N)
	for i,_ := range ret {
		ret[i] = make([]int, in.N)
	}
	for _,r := range in.Data {
		ret[r[0]][r[1]] = r[2]
		ret[r[1]][r[0]] = r[2]
	}
	return ret
}


func MST(in *MSTInput) []data.Edge {
	mat := HandleMSTInput(in)
	uf := union_find.UnionFind{}
	uf.Init(len(mat))
	ary := make([]min_heap.Comparable, 0)
	for i,v := range mat[0] {
		if v != 0 {
			ary = append(ary, data.Edge{
				0,
				i,
				v,
			})
		}
	}
	mh := min_heap.MinHeap{}
	mh.BuildMinHeap(ary)
	ans := make([]data.Edge, 0)
	// assume connected graph
	for !mh.Empty() {
		e := mh.Pop().(data.Edge)
		if uf.IsUnion(e.From, e.To) {
			continue
		}
		uf.Union(e.From, e.To)
		for i,v := range mat[e.To] {
			if v != 0 && !uf.IsUnion(e.To,i) {
				mh.Insert(data.Edge{
					From: e.To,
					To:   i,
					W:    v,
				})
			}
		}
		ans = append(ans, e)
	}
	return ans
}