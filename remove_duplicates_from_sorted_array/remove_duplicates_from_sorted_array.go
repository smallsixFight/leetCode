package remove_duplicates_from_sorted_array

func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	l := 1
	for p1, p2 := 0, 1; p2 < len(nums); p2++ {
		if nums[p1] == nums[p2] {
			continue
		}
		p1 += 1
		nums[p1] = nums[p2]
		l += 1
	}
	return l
}
