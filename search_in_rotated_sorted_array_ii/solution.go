package search_in_rotated_sorted_array_ii

func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	if nums[0] == target || nums[len(nums)-1] == target {
		return true
	}
	low, high := 0, len(nums)-1
	p := findRotatePosition(nums, target, 0, len(nums)-1)
	if nums[p] == target {
		return true
	}
	if target < nums[0] {
		low = p
	} else {
		low = 1
		if p != 0 {
			high = p - 1
		}
	}
	for low <= high {
		mid := (low + high) >> 1
		if nums[mid] == target {
			return true
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}

func findRotatePosition(nums []int, target, low, high int) int {
	for low <= high {
		mid := (low + high) >> 1
		if nums[mid] == target {
			return mid
		}
		if nums[0] < nums[mid] && mid+1 < len(nums) {
			if nums[0] >= nums[mid+1] {
				return mid + 1
			} else {
				low = mid + 1
			}
		} else if nums[0] > nums[mid] {
			high = mid - 1
		} else if nums[0] == nums[mid] {
			if r := findRotatePosition(nums, target, low, mid-1); r > 0 {
				return r
			}
			return findRotatePosition(nums, target, mid+1, high)
		} else {
			return 0
		}
	}
	return 0
}
