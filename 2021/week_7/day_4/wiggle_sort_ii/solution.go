package sort_list

import (
	"sort"
)

/*
324
来源：[leetCode](https://leetcode-cn.com/)
题目：[摆动排序 II](https://leetcode-cn.com/problems/wiggle-sort-ii/)
标签：`排序`

思路：先排序，然后拆分成前后两个数组，分别对两个数组进行逆序，接着将较大的数组进行插入到较小的数组中。


官网运行结果记录
执行用时：24ms/32ms
内存消耗：6.3MB

*/

func wiggleSort(nums []int) {
	sort.Ints(nums)
	temp, l := make([]int, len(nums)), len(nums)-1
	copy(temp, nums)
	for i := 1; i < len(nums); i += 2 {
		nums[i] = temp[l]
		l--
	}
	for i := 0; i < len(nums); i += 2 {
		nums[i] = temp[l]
		l--
	}
}
