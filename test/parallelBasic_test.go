package test

import (
	"algo/parallel"
	"testhelper"
	"testing"
)

func TestReduce(t *testing.T) {
	n := 10000
	ret := parallel.Reduce(testhelper.BigInput(n), testhelper.Plus, 0)
	testhelper.AssertEqual(t, (n-1)*n/2, ret)
}

func TestScan(t *testing.T) {
	ary := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ans := make([]int, len(ary))
	ans[0] = 1
	for i := 1; i < len(ary); i++ {
		ans[i] = ary[i] + ans[i-1]
	}
	ret := parallel.Scan(ary, testhelper.Plus, 0)
	for i := 0; i < len(ary); i++ {
		testhelper.AssertEqual(t, ans[i], ret[i])
	}
}

func TestScan2(t *testing.T) {
	ary := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ans := make([]int, len(ary))
	ans[0] = 1
	for i := 1; i < len(ary); i++ {
		ans[i] = ary[i] + ans[i-1]
	}
	ret := parallel.Scan2(ary, testhelper.Plus, 0)
	for i := 0; i < len(ary); i++ {
		testhelper.AssertEqual(t, ans[i], ret[i])
	}
	ary = testhelper.BigInput(5000)
	ans = scan(ary, testhelper.Plus)
	ret = parallel.Scan2(ary, testhelper.Plus, 0)
	for i := 0; i < len(ary); i++ {
		testhelper.AssertEqual(t, ans[i], ret[i])
	}
}

func TestScan3(t *testing.T) {
	ary := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ans := make([]int, len(ary))
	ans[0] = 1
	for i := 1; i < len(ary); i++ {
		ans[i] = ary[i] + ans[i-1]
	}
	ret := parallel.Scan3(ary, testhelper.Plus, 0, 4)
	for i := 0; i < len(ary); i++ {
		testhelper.AssertEqual(t, ans[i], ret[i])
	}
	ary = testhelper.BigInput(5000)
	ans = scan(ary, testhelper.Plus)
	ret = parallel.Scan3(ary, testhelper.Plus, 0, 4)
	for i := 0; i < len(ary); i++ {
		testhelper.AssertEqual(t, ans[i], ret[i])
	}
}

func TestScan4(t *testing.T) {
	ary := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ans := make([]int, len(ary))
	ans[0] = 1
	for i := 1; i < len(ary); i++ {
		ans[i] = ary[i] + ans[i-1]
	}
	ret := parallel.Scan4(ary, testhelper.Plus, 0, 4)
	for i := 0; i < len(ary); i++ {
		testhelper.AssertEqual(t, ans[i], ret[i])
	}
	ary = testhelper.BigInput(5000)
	ans = scan(ary, testhelper.Plus)
	ret = parallel.Scan4(ary, testhelper.Plus, 0, 4)
	for i := 0; i < len(ary); i++ {
		testhelper.AssertEqual(t, ans[i], ret[i])
	}
}

func scan(in []int, f func(int, int) int) []int {
	ret := make([]int, len(in))
	copy(ret, in)
	for i := 1; i < len(ret); i++ {
		ret[i] = f(ret[i-1], ret[i])
	}
	return ret
}
