package test

import (
	"algo/mst"
	"data"
	"testhelper"
	"testing"
)

func generateInput(s int) *mst.MSTInput {
	switch s {
	case 0:
		return &mst.MSTInput{
			N: 4,
			Data: [][]int{
				{0,1,2},
				{1,2,3},
				{2,3,4},
				{0,3,1},
			},
		}
	case 1:
		return &mst.MSTInput{
			N:    5,
			Data: [][]int{
				{0,1,5},
				{0,4,1},
				{0,3,4},
				{4,2,2},
				{4,3,3},
				{1,2,6},
				{2,3,7},
			},
		}
	}
	return nil
}

func TestInputHandling(t *testing.T) {
	in := generateInput(0)
	out := mst.HandleMSTInput(in)
	comp := [][]int{
		{0,2,0,1},
		{2,0,3,0},
		{0,3,0,4},
		{1,0,4,0},
	}
	for i,row := range out {
		for j,v := range row {
			if v != comp[i][j] {
				t.Errorf("Position (%d,%d) should be %d, got %d", i, j, comp[i][j], v)
			}
		}
	}
}

func TestMST(t *testing.T) {
	// simple case
	in := generateInput(0)
	out := mst.MST(in)
	testhelper.AssertEqual(t, 6, data.CalcPathLen(out))
	// complex case
	in = generateInput(1)
	out = mst.MST(in)
	testhelper.AssertEqual(t, 11, data.CalcPathLen(out))
}