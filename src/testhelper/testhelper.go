/*
 the code of this file cited from link below
版权声明：本文为CSDN博主「Cocolada」的原创文章，遵循 CC 4.0 BY-SA 版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/qq_30504647/article/details/83064574
 */

package testhelper

type testcommon interface {
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fail()
	FailNow()
	Failed() bool
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Helper()
	Log(args ...interface{})
	Logf(format string, args ...interface{})
	Name() string
	Skip(args ...interface{})
	SkipNow()
	Skipf(format string, args ...interface{})
	Skipped() bool
}

func AssertEqual(t testcommon, a, b interface{}) {
	t.Helper()
	if a != b {
		t.Errorf("Not Equal. %d %d", a, b)
	}
}

func AssertTrue(t testcommon, a bool) {
	t.Helper()
	if !a {
		t.Errorf("Not True %t", a)
	}
}

func AssertFalse(t testcommon, a bool) {
	t.Helper()
	if a {
		t.Errorf("Not True %t", a)
	}
}

func AssertNil(t testcommon, a interface{}) {
	t.Helper()
	if a != nil {
		t.Error("Not Nil")
	}
}

func AssertNotNil(t testcommon, a interface{}) {
	t.Helper()
	if a == nil {
		t.Error("Is Nil")
	}
}