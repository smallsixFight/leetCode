package gray_code

import "math"

func grayCode(n int) []int {
	l := int(math.Pow(2, float64(n)))
	res := make([]int, l)
	for i := 0; i < l; i++ {
		res[i] = i ^ (i >> 1)
	}
	return res
}
