package search_in_rotated_sorted_array

func Search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	//t1 := nums[0]
	low, high := 0, len(nums)-1
	minIdx := 0
	for low <= high {
		mid := (low + high) >> 1
		if nums[0] <= nums[mid] && mid+1 < len(nums) {
			if nums[0] > nums[mid+1] {
				minIdx = mid + 1
				break
			} else {
				low = mid + 1
			}
		} else if nums[0] > nums[mid] {
			high = mid - 1
		} else {
			break
		}
	}
	if target == nums[0] {
		return 0
	} else if target < nums[0] {
		low = minIdx
		high = len(nums) - 1
	} else {
		low = 1
		if minIdx != 0 { // 没有旋转的话 high 值应该 nums.length -1
			high = minIdx - 1
		} else {
			high = len(nums) - 1
		}
	}
	for low <= high {
		mid := (low + high) >> 1
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
