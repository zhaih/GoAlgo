package mst

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