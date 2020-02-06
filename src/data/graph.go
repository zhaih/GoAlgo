package data

type Edge struct {
	From, To, W int
}

func (e Edge) Less(e2 interface{}) bool {
	E := e2.(Edge)
	return e.W < E.W
}

func CalcPathLen(es []Edge) int {
	s := 0
	for _, e := range es {
		s += e.W
	}
	return s
}

func Edge2Matrix(es []Edge, n int, direct bool) [][]int {
	ret := make([][]int, n)
	for i := 0; i < n; i++ {
		ret[i] = make([]int, n)
	}
	for _, e := range es {
		ret[e.From][e.To] = e.W
		if !direct {
			ret[e.To][e.From] = e.W
		}
	}
	return ret
}

func Edge2Adj(es []Edge, n int, direct bool) [][]Edge {
	ret := make([][]Edge, n)
	for _, e := range es {
		if ret[e.From] == nil {
			ret[e.From] = make([]Edge, 0)
		}
		ret[e.From] = append(ret[e.From], Edge{
			From: e.From,
			To:   e.To,
			W:    e.W,
		})
		if !direct {
			if ret[e.To] == nil {
				ret[e.To] = make([]Edge, 0)
			}
			ret[e.To] = append(ret[e.To], Edge{
				From: e.To,
				To:   e.From,
				W:    e.W,
			})
		}
	}
	return ret
}

func GenBinaryTree(n int) []Edge {
	ret := make([]Edge, 0)
	for i := 0; i < n; i++ {
		l := 2*i + 1
		r := 2*i + 2
		if l < n {
			ret = append(ret, Edge{
				From: i,
				To:   l,
				W:    1,
			})
		}
		if r < n {
			ret = append(ret, Edge{
				From: i,
				To:   r,
				W:    1,
			})
		}
	}
	return ret
}
