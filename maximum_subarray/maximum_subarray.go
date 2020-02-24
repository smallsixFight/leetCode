package maximum_subarray

func MaxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	p1, p2 := 0, 1
	maxSum := nums[0]
	temp := maxSum
	for p2 < len(nums) {
		if nums[p1] < 0 && nums[p2] > nums[p1] {
			p1 = p2
			temp = nums[p1]
		} else {
			temp += nums[p2]
		}
		if temp > maxSum {
			maxSum = temp
		}
		p2++
		if temp < 0 { // temp 小于 0 处理
			p1 = p2
			temp = 0
		}
	}
	return maxSum
}
