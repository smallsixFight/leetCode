package sort_colors

func SortColors(nums []int) {
	if len(nums) < 2 {
		return
	}
	p1, p2 := 0, len(nums)
	for i := 0; i < p2 && p1 < p2; {
		if nums[i] == 0 && i != p1 {
			nums[p1], nums[i] = nums[i], nums[p1]
			p1++
		} else if nums[i] == 2 {
			nums[p2-1], nums[i] = nums[i], nums[p2-1]
			if i == p1 && nums[i] == 0 {
				i++
				p1++
			}
			p2--
		} else {
			i++
		}
	}
}
