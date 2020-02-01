package data

type GNode struct {
	Id int
	Weight int
	Edges []GEdge
}

type GEdge struct {
	A,B *GNode
	Weight int
}