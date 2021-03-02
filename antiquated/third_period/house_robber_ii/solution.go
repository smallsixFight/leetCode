package house_robber_ii

func Rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	return max(traverse(0, len(nums)-2, nums), traverse(1, len(nums)-1, nums))
}

func traverse(a, b int, nums []int) int {
	dp := make([]int, b-a+1)
	dp[0] = nums[a]
	dp[1] = max(nums[a], nums[a+1])
	for i := 2; i <= b-a; i++ {
		dp[i] = max(nums[a+i]+dp[i-2], dp[i-1])
	}
	return dp[b-a]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
