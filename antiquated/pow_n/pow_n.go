package pow_n

// 快速幂
func MyPowTwo(x float64, n int) float64 {
	res := 1.0
	if n < 0 {
		x = 1 / x
	}
	for n != 0 {
		if n%2 != 0 {
			res = res * x
		}
		x = x * x
		n = n / 2
	}
	return res
}

func MyPow(x float64, n int) float64 {
	if n == 0 { // 任何数的 0 次 幂都为 1
		return 1
	}
	pre := MyPow(x, n/2)
	if n%2 == 0 {
		return pre * pre
	}
	if n > 0 {
		return pre * pre * x
	}
	return pre * pre / x // n 小于 0
}
