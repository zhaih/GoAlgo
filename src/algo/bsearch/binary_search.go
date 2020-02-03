package bsearch

func Bsearch(t int, ary []int) int {
	l, u := 0, len(ary)
	for l < u {
		m := (l + u) / 2
		if ary[m] == t {
			return m
		}
		if ary[m] < t {
			l = m + 1
		} else {
			u = m
		}
	}
	return -1
}

// similar to c/c++ function upper_bound
// find a point x where ary[x] > t and ary[x-1] <= t
func UpperBound(t int, ary []int) int {
	l, u := 0, len(ary)
	for l < u {
		m := (l + u) / 2
		if ary[m] <= t {
			l = m + 1
		} else if m == l || ary[m-1] > t {
			u = m
		} else {
			return m
		}
	}
	return l
}

// similar to c/c++ lower_bound
// find a point x where ary[x]>=t and ary[x-1]<t
func LowerBound(t int, ary []int) int {
	l, u := 0, len(ary)
	for l < u {
		m := (l + u) / 2
		if ary[m] < t {
			l = m + 1
		} else if m == l || ary[m-1] >= t {
			u = m
		} else {
			return m
		}
	}
	return l
}
