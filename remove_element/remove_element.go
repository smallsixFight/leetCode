package remove_element

func RemoveElement(nums []int, val int) int {
	p1, p2 := 0, 0
	l := 0
	for p2 < len(nums) {
		if nums[p1] == val {
			if nums[p2] != val {
				nums[p1], nums[p2] = nums[p2], nums[p1]
				l += 1
				p1 += 1
			}
			p2 += 1
		} else {
			l += 1
			p1 += 1
			if nums[p2] == val || p1 >= p2 {
				p2 += 1
			}
		}
	}
	return l
}
