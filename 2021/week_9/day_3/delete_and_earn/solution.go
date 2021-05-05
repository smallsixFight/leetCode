package same_tree

import "sort"

/*
740 每日一题
来源：[leetCode](https://leetcode-cn.com/)
题目：[删除并获得点数](https://leetcode-cn.com/problems/delete-and-earn/)
标签：`动态规划`

简单思路：先对 `nums` 进行排序，把每个相同的整数看作一个整体，每个相同的数字计算累计和，再由于题目要求 `nums[i]` 不能与 `nums[i]-1` 和 `nums[i]+1` ，不过由于已排序，所以只需考虑 `nums[i]-1` 的情况，
那么就想到把 `nums[i] - nums[i-1] > 1` 作为一个临界进行分割，分别计算每个部分的最大值，然后叠加作为结果集。

官网运行结果记录
执行用时：4ms
内存消耗：3MB

*/

func deleteAndEarn(nums []int) int {
	sort.Ints(nums)
	res, sums := 0, []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			sums[len(sums)-1] += nums[i]
		} else if nums[i] == nums[i-1]+1 {
			sums = append(sums, nums[i])
		} else {
			res += getSum(sums)
			sums = []int{nums[i]}
		}
	}
	res += getSum(sums)
	return res
}

func getSum(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	a, b := arr[0], max(arr[0], arr[1])
	for i := 2; i < len(arr); i++ {
		a, b = b, max(a+arr[i], b)
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
