package search_insert_position

func SearchInsert(nums []int, target int) int {
	low, high := 0, len(nums)-1
	if target < nums[low] {
		return 0
	}
	if target > nums[high] {
		return len(nums)
	}
	for low <= high {
		mid := (low + high) >> 1
		if target < nums[mid] {
			high = mid - 1
		} else if target > nums[mid] {
			low = mid + 1
		} else {
			return mid
		}
	}
	if target > nums[low] {
		low++
	}
	return low
}
