package test

import (
	"algo/qsearch"
	"testing"
)
import "testhelper"

func TestQsearch(t *testing.T) {
	ary := []int{1,6,4,5,3,2,7,8,9}
	testhelper.AssertEqual(t, 1, qsearch.Qsearch(1,0,ary))
	testhelper.AssertEqual(t, 3, qsearch.Qsearch(3,0,ary))
	testhelper.AssertEqual(t, 6, qsearch.Qsearch(6,0,ary))
	testhelper.AssertEqual(t, 9, qsearch.Qsearch(9,0,ary))
}
