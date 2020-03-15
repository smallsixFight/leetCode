package find_first_and_last_position_of_element_in_sorted_array

func SearchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	if len(nums) == 0 {
		return res
	}
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (low + high) >> 1
		if nums[mid] == target {
			if mid == 0 || (mid > 0 && nums[mid-1] != target) {
				res[0] = mid
				break
			} else {
				high = mid - 1
			}
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if res[0] == -1 {
		return res
	}
	if res[0] == len(nums)-1 {
		res[1] = res[0]
		return res
	}
	low, high = res[0], len(nums)-1
	for low <= high {
		mid := (low + high) >> 1
		if nums[mid] == target {
			if mid == len(nums)-1 || (mid+1 < len(nums) && nums[mid+1] != target) {
				res[1] = mid
				break
			} else {
				low = mid + 1
			}
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return res
}
