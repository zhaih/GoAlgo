package bench

import (
	"algo/parallel"
	"testhelper"
	"testing"
)

func BenchmarkSeqScan(b *testing.B) {
	input := testhelper.BigInput(1000000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		scan(input, testhelper.Plus)
	}
}

func BenchmarkParaScan(b *testing.B) {
	input := testhelper.BigInput(1000000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		parallel.Scan(input, testhelper.Plus, 0)
	}
}

func BenchmarkParaScan2(b *testing.B) {
	input := testhelper.BigInput(1000000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		parallel.Scan2(input, testhelper.Plus, 0)
	}
}

func BenchmarkParaScan3(b *testing.B) {
	input := testhelper.BigInput(1000000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		parallel.Scan3(input, testhelper.Plus, 0, 4)
	}
}

func BenchmarkParaScan4(b *testing.B) {
	input := testhelper.BigInput(1000000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		parallel.Scan4(input, testhelper.Plus, 0, 4)
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
