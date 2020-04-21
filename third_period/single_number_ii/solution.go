package single_number_ii

// TODO
func SingleNumbersTwo(nums []int) int {
	a, b := 0, 0
	for _, v := range nums {
		b = (b ^ v) & (^a)
		a = (a ^ v) & (^b)
	}
	return b
}

func SingleNumbers(nums []int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	for k := range m {
		if m[k] == 1 {
			return k
		}
	}
	return -1
}
