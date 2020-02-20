package test

import (
	"data"
	"testhelper"
	"testing"
)

func TestSetInit(t *testing.T) {
	is := data.IntSet{}
	is.Init([]int{1, 2, 3, 4})
	testhelper.AssertTrue(t, is.Contains(1))
	testhelper.AssertTrue(t, is.Contains(2))
	testhelper.AssertTrue(t, is.Contains(3))
	testhelper.AssertTrue(t, is.Contains(4))
	testhelper.AssertFalse(t, is.Contains(5))
}

func TestSetAdd(t *testing.T) {
	is := data.IntSet{}
	is.Init([]int{1, 2, 3, 4})
	is.Add(5)
	testhelper.AssertTrue(t, is.Contains(5))
}

func TestSetRemove(t *testing.T) {
	is := data.IntSet{}
	is.Init([]int{1, 2, 3, 4})
	is.Remove(2)
	testhelper.AssertFalse(t, is.Contains(2))
	is.Remove(5)
	testhelper.AssertFalse(t, is.Contains(5))
}

func TestSetSize(t *testing.T) {
	is := data.IntSet{}
	is.Init([]int{1, 2, 3, 4})
	testhelper.AssertEqual(t, is.Size(), 4)
	is.Add(6)
	is.Add(7)
	is.Add(1000)
	testhelper.AssertEqual(t, is.Size(), 7)
	is.Add(1000)
	is.Add(8)
	is.Add(6)
	testhelper.AssertEqual(t, is.Size(), 8)
}

func TestSetIterate(t *testing.T) {
	is := data.IntSet{}
	is.Init([]int{1, 2, 3, 4})
	is.Add(5)
	testhelper.AssertEqual(t, 1, is.Get())
	testhelper.AssertTrue(t, is.Next())
	testhelper.AssertEqual(t, 2, is.Get())
	testhelper.AssertTrue(t, is.Next())
	testhelper.AssertEqual(t, 3, is.Get())
	testhelper.AssertTrue(t, is.Next())
	testhelper.AssertEqual(t, 4, is.Get())
	testhelper.AssertTrue(t, is.Next())
	testhelper.AssertEqual(t, 5, is.Get())
	testhelper.AssertFalse(t, is.Next())
}

func TestSetInsert(t *testing.T) {
	is := data.IntSet{}
	is.Init([]int{1, 2, 3, 4})
	is.Next()
	is.Next()
	is.Insert(8)
	testhelper.AssertTrue(t, is.Contains(8))
	for i := 0; i < is.Size()-1; i++ {
		is.Next()
	}
	testhelper.AssertEqual(t, is.Get(), 8)
}

func TestSetDelete(t *testing.T) {
	is := data.IntSet{}
	is.Init([]int{1, 2, 3, 4})
	is.Next()
	is.Next()
	is.Delete()
	testhelper.AssertFalse(t, is.Contains(3))
	testhelper.AssertEqual(t, is.Get(), 4)
}
