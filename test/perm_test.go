package test

import (
	"algo/back_track"
	"testhelper"
	"testing"
)

func TestPermutation(t *testing.T) {
	ret := back_track.Permutation([]int{1, 2, 3})
	ans := [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 3, 1},
		{2, 1, 3},
		{3, 1, 2},
		{3, 2, 1},
	}
	for i, a := range ret {
		for j := 0; j < 3; j++ {
			testhelper.AssertEqual(t, ans[i][j], a[j])
		}
	}
}
