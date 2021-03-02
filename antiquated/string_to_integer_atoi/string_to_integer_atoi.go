package string_to_integer_atoi

import "math"

func MyAtoi(str string) int {
	p := 0
	flag := uint8('+')
	for p < len(str) {
		if str[p] != ' ' && str[p] != '+' && str[p] != '-' &&
			(str[p] < 47 || str[p] > 58) {
			return 0
		} else if str[p] == ' ' {
			p++
			continue
		} else {
			flag = str[p]
			break
		}
	}
	if flag == '+' || flag == '-' {
		p++
	} else {
		flag = '+'
	}
	res := 0
	for p < len(str) {
		if str[p] < 47 || str[p] > 58 {
			break
		}
		val := str[p] - '0'
		if flag == '+' {
			if res > math.MaxInt32/10 || (res == math.MaxInt32/10 && val > 7) {
				return math.MaxInt32
			}
			res = res*10 + int(val)
		} else {
			if res < math.MinInt32/10 || (res == math.MinInt32/10 && val > 8) {
				return math.MinInt32
			}
			res = res*10 - int(val)
		}
		p++
	}
	return res
}
