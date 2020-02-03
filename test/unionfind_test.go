package test

import (
	"algo/union_find"
	"testing"
)
import "testhelper"

func TestInit(t *testing.T) {
	uf := union_find.UnionFind{}
	uf.Init(5)
	if len(uf.Status()) != 5 {
		t.Errorf("Init length error: should be %d, got %d", 5, len(uf.Status()))
	}
	for i, n := range uf.Status() {
		if n != -1 {
			t.Errorf("Init value error at index %d", i)
		}
	}
}

func TestFind(t *testing.T) {
	uf := union_find.UnionFind{}
	uf.Init(5)
	for i := 0; i < 5; i++ {
		v := uf.Find(i)
		if v != i {
			t.Errorf("Find root value error, expect %d, got %d", i, v)
		}
	}
	uf.Status()[1] = 0
	uf.Status()[2] = 1
	testhelper.AssertEqual(t, uf.Find(1), 0)
	testhelper.AssertEqual(t, uf.Find(2), 0)
	testhelper.AssertEqual(t, uf.Status()[2], 0)
}

func TestUnion(t *testing.T) {
	uf := union_find.UnionFind{}
	uf.Init(10)
	// Union 0,2,4,6,8
	for i := 0; i < 8; i += 2 {
		testhelper.AssertEqual(t, 0, uf.Union(i, i+2))
	}
	// Union 1,3,5,7,9
	for i := 1; i < 9; i += 2 {
		testhelper.AssertEqual(t, 1, uf.Union(i, i+2))
	}
	testhelper.AssertEqual(t, 0, uf.Union(3, 6))
	testhelper.AssertEqual(t, 0, uf.Find(9))
}
