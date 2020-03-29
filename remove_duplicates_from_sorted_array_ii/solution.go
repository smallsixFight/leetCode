package remove_duplicates_from_sorted_array_ii

func RemoveDuplicatesTwo(nums []int) int {
	k := 1 // [0,k] 存放需要保留的元素。初始化k=1，因为[0,1]是肯定需要放入数组中的
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[k-1] {
			k++
			nums[k] = nums[i]
		}
	}
	return k + 1
}

func RemoveDuplicates(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}
	p1, p2 := 0, 0
	count := 1
	for p2 < len(nums) {
		if p2 > 0 {
			if nums[p2] == nums[p2-1] {
				count++
			} else {
				count = 1
			}
		}
		if count < 3 {
			nums[p1] = nums[p2]
			p1++
		}
		p2++
	}
	return p1
}
