package sqrt_x

import "math"

func MySqrt2(x int) int {
	if x < 2 {
		return x
	}
	x1 := float64(x)
	x2 := (x1 + float64(x)/x1) / 2
	for math.Abs(float64(x1-x2)) >= 1 {
		x1 = x2
		x2 = (x1 + float64(x)/x1) / 2
	}
	return int(x2)
}

func MySqrt(x int) int {
	res := 1
	for low, high := 0, x; low <= high; {
		mid := (low + high) >> 1
		temp := mid * mid
		if temp == x {
			res = mid
			break
		} else if temp < x {
			res = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return res
}
