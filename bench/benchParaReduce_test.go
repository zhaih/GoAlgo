package bench

import (
	"algo/parallel"
	"testhelper"
	"testing"
)

func BenchmarkSeqReduce(b *testing.B) {
	input := testhelper.BigInput(1000000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reduce(input, testhelper.Plus)
	}
}

func BenchmarkParaReduce(b *testing.B) {
	input := testhelper.BigInput(1000000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		parallel.Reduce(input, testhelper.Plus, 0)
	}
}

func reduce(in []int, f func(int, int) int) int {
	ans := in[0]
	for i := 1; i < len(in); i++ {
		ans = f(ans, in[i])
	}
	return ans
}
