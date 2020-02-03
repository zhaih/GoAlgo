package test

import (
	"algo/bsearch"
	"testing"
)
import "testhelper"

func TestBsearch(t *testing.T) {
	ary := make([]int, 11)
	for i := 0; i < 10; i++ {
		ary[i] = i + 1
	}
	ary[10] = 12
	for i := 0; i < 10; i++ {
		testhelper.AssertEqual(t, i, bsearch.Bsearch(i+1, ary))
	}
	testhelper.AssertEqual(t, -1, bsearch.Bsearch(100, ary))
	testhelper.AssertEqual(t, -1, bsearch.Bsearch(0, ary))
	testhelper.AssertEqual(t, -1, bsearch.Bsearch(11, ary))
}

func TestLowerBound(t *testing.T) {
	ary := make([]int, 11)
	for i := 0; i < 10; i++ {
		ary[i] = i + 1
	}
	ary[10] = 10
	for i := 0; i < 10; i++ {
		testhelper.AssertEqual(t, i, bsearch.LowerBound(i+1, ary))
	}
	testhelper.AssertEqual(t, 0, bsearch.LowerBound(0, ary))
	testhelper.AssertEqual(t, 11, bsearch.LowerBound(100, ary))
}

func TestUpperBound(t *testing.T) {
	ary := make([]int, 11)
	for i := 0; i < 10; i++ {
		ary[i] = i + 1
	}
	ary[10] = 10
	for i := 0; i < 10; i++ {
		testhelper.AssertEqual(t, i, bsearch.UpperBound(i, ary))
	}
	testhelper.AssertEqual(t, 11, bsearch.UpperBound(10, ary))
	testhelper.AssertEqual(t, 11, bsearch.UpperBound(100, ary))
}
