package maximum_product_subarray

func MaxProductTwo(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	res, preMax, preMin := nums[0], 1, 1
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			preMax, preMin = preMin, preMax
		}
		preMax = max(preMax*nums[i], nums[i])
		preMin = min(preMin*nums[i], nums[i])
		res = max(preMax, res)
	}
	return res
}

func MaxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([][]int, len(nums))
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, len(nums))
		dp[i][i] = nums[i]
		if nums[i] > res {
			res = nums[i]
		}
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			dp[i][j] = dp[i][j-1] * dp[j][j]
			if dp[i][j] > res {
				res = dp[i][j]
			}
		}
	}
	return res
}
