package sort_list

import "sort"

/*
349
来源：[leetCode](https://leetcode-cn.com/)
题目：[两个数组的交集](https://leetcode-cn.com/problems/intersection-of-two-arrays/)
标签：`排序` `哈希表` `双指针` `二分查找`

思路：先对两个数组进行排序，然后遍历个数较少的数组，遍历时如果 `nums[i] == nums[i-1]`，那么跳过；否则在另外一个数组中进行查找，存在加入结果集。

官网运行结果记录
执行用时：4ms
内存消耗：2.7MB

*/

func intersection(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	sort.Ints(nums1)
	sort.Ints(nums2)
	if len(nums2) < len(nums1) {
		nums1, nums2 = nums2, nums1
	}
	for i := range nums1 {
		if i == 0 || nums1[i] != nums1[i-1] {
			if isExist(nums2, nums1[i]) {
				res = append(res, nums1[i])
			}
		}
	}
	return res
}

func isExist(arr []int, target int) bool {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := (low + high) >> 1
		if target < arr[mid] {
			high = mid - 1
		} else if target > arr[mid] {
			low = mid + 1
		} else {
			return true
		}
	}
	return false
}
