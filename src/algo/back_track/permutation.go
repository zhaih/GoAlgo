package back_track

import "data"

func Permutation(ary []int) [][]int {
	is := data.IntSet{}
	is.Init(ary)
	ret := make([][]int, 0)

	var next func([]int)
	next = func(r []int) {
		if is.Size() == 1 {
			r = append(r, is.Get())
			ret = append(ret, r)
			return
		}
		s := is.Size()
		for i := 0; i < s; i++ {
			x := is.Get()
			is.Delete()
			entry := make([]int, len(r)+1)
			copy(entry, r)
			entry[len(r)] = x
			next(entry)
			is.Insert(x)
		}
	}
	next([]int{})
	return ret
}
