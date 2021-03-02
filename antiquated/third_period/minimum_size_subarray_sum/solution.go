package minimum_size_subarray_sum

func MinSubArrayLen(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res, curSum := 0, nums[0]
	if curSum >= s {
		return 1
	}
	p1, p2 := 0, 1
	for p1 < len(nums) && p2 <= len(nums) {
		if curSum >= s {
			if res == 0 || p2-p1 < res {
				res = p2 - p1
			}
			curSum -= nums[p1]
			p1++
		} else {
			if p2 == len(nums) {
				break
			}
			curSum += nums[p2]
			p2++
		}
	}
	return res
}
