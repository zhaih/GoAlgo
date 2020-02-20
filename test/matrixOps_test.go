package test

import (
	"algo/math"
	"testhelper"
	"testing"
)

func TestMatMulSlow(t *testing.T) {
	ma := testhelper.GenMat(4, 6, 2)
	mb := testhelper.GenEyeMat(6)
	out := math.MatmulSlow(ma, mb)

	for i, r := range ma {
		for j, v := range r {
			testhelper.AssertEqual(t, v, out[i][j])
		}
	}
}

func TestMatMulTiling(t *testing.T) {
	ma := testhelper.GenMat(4, 6, 2)
	mb := testhelper.GenEyeMat(6)
	out := math.MatmulTiling(ma, mb, 2)

	for i, r := range ma {
		for j, v := range r {
			testhelper.AssertEqual(t, v, out[i][j])
		}
	}

	ma2 := testhelper.GenMat(4, 4, 1)
	out2 := math.MatmulTiling(ma2, ma2, 2)
	for i, r := range ma2 {
		for j, v := range r {
			testhelper.AssertEqual(t, 4*v, out2[i][j])
		}
	}
}
