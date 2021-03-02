package palindrome_number

import "fmt"

func FirstMethod(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	str := fmt.Sprintf("%d", x)
	i, j := 0, len(str)-1
	for i < j {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func SecondMethod(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	reverseVal := 0
	for x > reverseVal {
		reverseVal = reverseVal*10 + x%10
		x /= 10
	}
	return reverseVal == x || reverseVal/10 == x
}
