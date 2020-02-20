package bench

import (
	"algo/math"
	"testhelper"
	"testing"
)

const test_size = 500

/*
 Test result when ma and mb are both 5000 x 5000 matrix
 goos: windows
 goarch: amd64
 BenchmarkMatMulSlow-8       	       1	2114916672000 ns/op
 BenchmarkMatMulTiling16-8   	       1	362979711200 ns/op
 BenchmarkMatMulTiling24-8   	       1	373598301500 ns/op
 BenchmarkMatMulTiling32-8   	       1	376010471000 ns/op
*/

func BenchmarkMatMulSlow(b *testing.B) {
	ma := testhelper.GenMat(test_size, test_size, 1)
	mb := testhelper.GenMat(test_size, test_size, 2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		math.MatmulSlow(ma, mb)
	}
}

func BenchmarkMatMulTiling16(b *testing.B) {
	ma := testhelper.GenMat(test_size, test_size, 1)
	mb := testhelper.GenMat(test_size, test_size, 2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		math.MatmulTiling(ma, mb, 16)
	}
}

func BenchmarkMatMulTiling24(b *testing.B) {
	ma := testhelper.GenMat(test_size, test_size, 1)
	mb := testhelper.GenMat(test_size, test_size, 2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		math.MatmulTiling(ma, mb, 24)
	}
}

func BenchmarkMatMulTiling32(b *testing.B) {
	ma := testhelper.GenMat(test_size, test_size, 1)
	mb := testhelper.GenMat(test_size, test_size, 2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		math.MatmulTiling(ma, mb, 32)
	}
}
