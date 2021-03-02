package plus_one

func PlusOne(digits []int) []int {
	i := len(digits) - 1
	for ; i >= 0; i-- {
		digits[i] += 1
		if digits[i] < 10 {
			break
		}
		digits[i] = 0
	}
	if i == -1 && digits[i+1] == 0 {
		arr := make([]int, len(digits)+1)
		arr[0] = 1
		for j := 0; j < len(digits); j++ {
			arr[j+1] = digits[j]
		}
		return arr
	} else {
		return digits
	}
}
