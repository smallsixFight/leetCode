package fraction_to_recurring_decimal

import (
	"strconv"
	"strings"
)

func FractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}
	if denominator == 0 {
		return ""
	}
	res := strings.Builder{}
	preNum := abs(numerator / denominator)
	if (numerator < 0 && denominator > 0) || (numerator > 0 && denominator < 0) {
		res.WriteByte('-')
	}
	res.WriteString(strconv.FormatInt(int64(preNum), 10))
	remainder := numerator % denominator
	if remainder == 0 {
		return res.String()
	}
	res.WriteByte('.')
	n1, n2 := 0, 0
	temp := denominator
	for denominator%10 == 0 && denominator >= 10 {
		n1++
		denominator /= 10
	}

	for numerator%10 == 0 && numerator >= 10 {
		n2++
		numerator /= 10

	}
	denominator = abs(denominator)
	remainder = abs(remainder)
	for n1 > n2 && remainder*10 < temp {
		res.WriteByte('0')
		temp /= 10
		n1--
	}

	other := make([]byte, 0)
	m := make(map[int][]int)
	record := make(map[int]int)
	idx := 0
	isZero := true
	for remainder > 0 {
		if remainder < denominator {
			remainder *= 10
		}
		num := remainder / denominator
		if num != 0 {
			isZero = false
		}
		addNum := 1
		if num > 10 {
			a := strconv.FormatInt(int64(num), 10)
			other = append(other, []byte(a)...)
			addNum = len(a)
		} else {
			other = append(other, byte(num)+'0')
		}
		remainder %= denominator
		record[idx] = remainder
		if remainder == 0 {
			break
		}
		if arr, ok := m[num]; ok {
			for i := len(arr) - 1; i >= 0 && !isZero; i-- {
				a := idx - arr[i]
				if record[arr[i]] == remainder && arr[i]-a > -1 && string(other[arr[i]:idx]) == string(other[arr[i]-a:arr[i]]) {
					temp := make([]byte, 0)
					temp = append(temp, other[:arr[i]-a]...)
					temp = append(temp, '(')
					temp = append(temp, other[arr[i]:idx]...)
					temp = append(temp, ')')
					other = temp
					return res.String() + string(other)
				}
			}
			m[num] = append(m[num], idx)
		} else {
			m[num] = []int{idx}
		}
		idx += addNum
	}
	return res.String() + string(other)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
