package divide_two_integers

import "math"

func Divide(dividend, divisor int) int {
	flag := false
	if dividend < 0 {
		dividend = -dividend
		flag = !flag
	}
	if divisor < 0 {
		divisor = -divisor
		flag = !flag
	}
	if dividend < divisor {
		return 0
	}
	res := 1
	temp := divisor << 1
	for temp < dividend {
		res <<= 1
		temp <<= 1
	}
	temp >>= 1
	if temp < dividend {
		res += Divide(dividend-temp, divisor)
	}
	if flag {
		res = -res
	}
	if res > math.MaxInt32 {
		res = math.MaxInt32
	}
	return res
}
