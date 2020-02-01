package test

import "testing"
import "testhelper"
import "algo/min_heap"

type tester int
func (i tester)Less(j interface{}) bool {
	J := int(j.(tester))
	I := int(i)
	return I<J
}

func TestMinHeap(t *testing.T) {
	a := []tester{1,3,4,7,8,2,10,100}
	sorted_a := []tester{1,2,3,4,7,8,10,100}
	ary := make([]min_heap.Comparable, len(a))
	for i:=0;i<len(a);i++ {
		ary[i] = a[i]
	}
	mh := min_heap.MinHeap{}
	mh.BuildMinHeap(ary)

	for i:=0;i<len(a);i++ {
		testhelper.AssertEqual(t, sorted_a[i], mh.Pop().(tester))
	}

	for i:=0;i<len(a);i++ {
		mh.Insert(a[i])
	}
	mh.Insert(tester(0))
	testhelper.AssertEqual(t,tester(0),mh.Pop().(tester))
	testhelper.AssertEqual(t, tester(1), mh.Pop().(tester))
}
