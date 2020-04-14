package first_missing_positive

func FirstMissingPositive(nums []int) int {
	hasOne := false
	l := len(nums)
	for i := 0; i < l; i++ {
		if nums[i] == 1 {
			hasOne = true
			break
		}
	}
	if !hasOne {
		return 1
	}
	if l == 1 {
		return 2
	}
	for i := 0; i < l; i++ {
		if nums[i] < 1 || nums[i] > l {
			nums[i] = 1
		}
	}
	getAbs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	for i := 0; i < l; i++ {
		idx := getAbs(nums[i])
		if idx < l {
			nums[idx] = -getAbs(nums[idx])
		} else {
			nums[0] = -getAbs(nums[0])
		}
	}
	for i := 1; i < l; i++ {
		if nums[i] > 0 {
			return i
		}
	}
	if nums[0] > 0 {
		return l
	}
	return l + 1
}
