package qsearch

import "math/rand"

// quick search
// Search n-th smallest element from ary, where n is the parameter rank
// Actually need copy to not mess up original data
// Could return index instead of data, but need copy mentioned above and
// maybe a map to remember original position
func Qsearch(rank,s int, ary []int) int {
	if rank == 1 {
		return ary[s]
	}
	ri := rand.Intn(len(ary) - s)
	p := ary[s+ri]
	ary[s+ri],ary[s] = ary[s],ary[s+ri]
	j := s + 1
	// partition to [p;>=p;<p]
	for i:=s+1;i<len(ary);i++ {
		if ary[i] >= p {
			if i != j {
				ary[i],ary[j] = ary[j],ary[i]
			}
			j++
		}
	}
	if len(ary)-j < rank {
		return Qsearch(rank+j-len(ary), s, ary[:j])
	} else {
		return Qsearch(rank, j, ary)
	}
}
