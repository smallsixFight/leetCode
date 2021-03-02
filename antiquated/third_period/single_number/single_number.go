package single_number

func SingleNumber(nums []int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	for k := range m {
		if m[k] == 1 {
			return k
		}
	}
	return 0
}
