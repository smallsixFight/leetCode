package sort_list

import "sort"

/*
164
来源：[leetCode](https://leetcode-cn.com/)
题目：[最大差距](https://leetcode-cn.com/problems/maximum-gap/)
标签：`排序`

简单思路：这道题就是要先排序，不过我所能掌握的排序方法并没有能够符合线性时间复杂度的，通过题解，学习到了一种叫 `基数排序` 的排序算法。

在学习理解的这个算法的时候再次感受他人的智慧。

官网运行结果记录
执行用时：4ms(maximumGap2，快排)/4ms(maximumGap，基数排序)
内存消耗：3.1MB(maximumGap2，快排)/3.7MB(maximumGap，基数排序)

*/

func maximumGap2(nums []int) int {
	sort.Ints(nums)
	res := 0
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] > res {
			res = nums[i] - nums[i-1]
		}
	}
	return res
}

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	maxVal := getMaxVal(nums)
	buckets := [10][]int{} // 这里只需要考虑非负整数
	mod, dev := 10, 1
	for maxVal >= 0 { // 最大需要处理的位数
		idx := 0
		for i := range buckets {
			for j := range buckets[i] {
				nums[idx] = buckets[i][j]
				idx++
			}
			buckets[i] = buckets[i][:0] // 	清空
		}
		if maxVal == 0 { // 0 还需要再从 buckets 取出一次，之后即可直接返回
			break
		}
		for i := range nums {
			buckets[nums[i]%mod/dev] = append(buckets[nums[i]%mod/dev], nums[i])
		}
		maxVal /= 10
		mod *= 10
		dev *= 10
	}
	res := nums[1] - nums[0]
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] > res {
			res = nums[i] - nums[i-1]
		}
	}
	return res
}

func getMaxVal(nums []int) int {
	val := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > val {
			val = nums[i]
		}
	}
	return val
}
