package find_peak_element

func FindPeakElement(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (low + high) >> 1
		if (mid == 0 && nums[0] > nums[1]) || (mid == len(nums)-1 && nums[len(nums)-1] > nums[len(nums)-2]) {
			return mid
		}
		if mid > 0 && mid < len(nums) {
			if nums[mid-1] < nums[mid] && nums[mid] > nums[mid+1] {
				return mid
			}
		}
		if mid == 0 || (nums[mid-1] < nums[mid] && nums[mid] < nums[mid+1]) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return 0
}
