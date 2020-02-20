package math

// const tiling_size = 32
// 32^2 = 1024, my cpu's l1 cache could store about 6400 int,
// need store 3 matrix at the same time, so it should be a reasonable number

// will not do error handling as it is not an actual ready-to-use function
func MatmulSlow(ma, mb [][]int) [][]int {
	if len(ma) == 0 || len(mb) == 0 || len(ma[0]) != len(mb) {
		return nil
	}

	ret := make([][]int, len(ma))
	for i := 0; i < len(ma); i++ {
		ret[i] = make([]int, len(mb[0]))
		for j := 0; j < len(mb[0]); j++ {
			for k := 0; k < len(mb); k++ {
				ret[i][j] += ma[i][k] * mb[k][j]
			}
		}
	}
	return ret
}

func MatmulTiling(ma, mb [][]int, t int) [][]int {
	if len(ma) == 0 || len(mb) == 0 || len(ma[0]) != len(mb) {
		return nil
	}

	ret := make([][]int, len(ma))
	for i := 0; i < len(ret); i++ {
		ret[i] = make([]int, len(mb[0]))
	}
	for i := 0; i < len(ret); i += t {
		for j := 0; j < len(mb[0]); j += t {
			for k := 0; k < len(mb); k += t {
				for x := 0; x < t && (i+x) < len(ma); x++ {
					for y := 0; y < t && (j+y) < len(mb[0]); y++ {
						for z := 0; z < t && (k+z) < len(mb); z++ {
							ret[i+x][j+y] += ma[i+x][k+z] * mb[k+z][j+y]
						}
					}
				}
			}
		}
	}
	return ret
}
