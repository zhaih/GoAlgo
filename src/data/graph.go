package data

type Edge struct {
	From, To int
	W        int
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
