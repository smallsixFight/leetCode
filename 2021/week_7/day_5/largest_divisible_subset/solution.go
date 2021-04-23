package largest_divisible_subset

import (
	"sort"
)

/*
368 每日一题
来源：[leetCode](https://leetcode-cn.com/)
题目：[最大整除子集](https://leetcode-cn.com/problems/largest-divisible-subset/)
标签：`数学` `动态规划`

思路：...

官网运行结果记录
执行用时：44ms
内存消耗：2.8MB

*/

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	maxSize, maxVal := 0, 0
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if dp[j] >= dp[i] && nums[i]%nums[j] == 0 {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > maxSize {
			maxSize, maxVal = dp[i], nums[i]
		}
	}
	if maxSize == 0 {
		return []int{nums[0]}
	}
	res := make([]int, 0)
	for i := len(nums) - 1; i >= 0 && maxSize > 0; i-- {
		if dp[i] == maxSize && maxVal%nums[i] == 0 {
			res = append(res, nums[i])
			maxVal = nums[i]
			maxSize--
		}
	}
	return res
}
