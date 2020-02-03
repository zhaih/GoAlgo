// Union find specially for integer
package union_find

type UnionFind struct {
	mem []int
}

func (uf *UnionFind) Init(n int) {
	uf.mem = make([]int, n)
	for i, _ := range uf.mem {
		uf.mem[i] = -1
	}
}

func (uf *UnionFind) Union(a, b int) int {
	sa := uf.Find(a)
	sb := uf.Find(b)
	if sa == sb {
		return sa
	}
	if sa < sb {
		uf.mem[sb] = sa
		return sa
	} else {
		uf.mem[sa] = sb
		return sb
	}
}

func (uf *UnionFind) Find(a int) int {
	if uf.mem[a] == -1 {
		return a
	}
	// to save a level of recursion and memory write
	if uf.mem[uf.mem[a]] == -1 {
		return uf.mem[a]
	}
	set := uf.Find(uf.mem[a])
	uf.mem[a] = set
	return set
}

func (uf *UnionFind) IsUnion(a, b int) bool {
	return uf.Find(a) == uf.Find(b)
}

// This function is purely for test purpose
func (uf *UnionFind) Status() []int {
	return uf.mem
}
