package reverse_integer

import (
	"math"
)

// 懒得引入自己写的 stack struct，用 slice 代替
func Reverse(x int) int {
	result := 0
	for x != 0 {
		val := x % 10
		x /= 10
		if (result > math.MaxInt32/10) || (result == math.MaxInt32/10 && val > 7) {
			return 0
		}
		if (result < math.MinInt32/10) || (result == math.MinInt32/10 && val < -8) {
			return 0
		}
		result = result*10 + val
	}
	return result
}
