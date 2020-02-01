package min_heap

type Comparable interface {
	Less(interface{}) bool
}

type MinHeap struct {
	ary []Comparable
	bound int
}

func heapify(ary []Comparable, root int) {
	left := 2*root + 1
	right := 2*root + 2
	mini := root
	if left < len(ary) && ary[left].Less(ary[mini]){
		mini = left
	}
	if right < len(ary) && ary[right].Less(ary[mini]){
		mini = right
	}
	if mini != root {
		ary[root], ary[mini] = ary[mini], ary[root]
		heapify(ary, mini)
	}
}

func (mh *MinHeap) BuildMinHeap(ary []Comparable) {
	for i:=len(ary)/2-1;i >= 0; i-- {
		heapify(ary, i)
	}
	mh.ary = ary
	mh.bound = len(ary)
}

func (mh *MinHeap) Pop() Comparable {
	ary := mh.ary
	mh.bound--
	ary[0], ary[mh.bound] = ary[mh.bound], ary[0]
	view := ary[:mh.bound]
	heapify(view, 0)
	return ary[mh.bound]
}

func (mh *MinHeap) Insert(c Comparable) {
	if mh.bound == len(mh.ary) {
		mh.ary = append(mh.ary, c)
	} else {
		mh.ary[mh.bound] = c
	}
	mh.bound++
	if mh.bound == 1 {
		return
	}
	view := mh.ary[:mh.bound]
	for b:=mh.bound/2-1;;b=(b-1)/2 {
		heapify(view, b)
		if b == 0 {
			break
		}
	}
}