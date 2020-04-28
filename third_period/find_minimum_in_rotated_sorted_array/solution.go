package find_minimum_in_rotated_sorted_array

func FindMin(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (low + high) >> 1
		if nums[mid] < nums[0] && nums[mid-1] >= nums[0] {
			return nums[mid]
		} else if nums[mid] < nums[0] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return nums[0]
}
