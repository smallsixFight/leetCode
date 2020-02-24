package maximum_subarray

func MaxSubArray2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	r := make([]int, len(nums))
	r[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if r[i-1] > 0 {
			r[i] = r[i-1] + nums[i]
		} else {
			r[i] = nums[i]
		}
	}
	res := r[0]
	for _, n := range r {
		if n > res {
			res = n
		}
	}
	return res
}

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
