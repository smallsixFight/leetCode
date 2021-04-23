package largest_divisible_subset

import "sort"

/*
350
来源：[leetCode](https://leetcode-cn.com/)
题目：[两个数组的交集 II](https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/)
标签：`排序` `哈希表` `双指针` `二分查找`

思路：分别遍历两个数组，使用两个 hashMap 分别记录两个数组出现个数字和出现个数，	然后遍历较小的 hashMap，如果这个数字在另外一个 hashMap 也出现了，去他们中最小出现个数，循环加入结果集。

思路二：先对两个数字进行排序，然后使用双指针法进行遍历比较。

官网运行结果记录
执行用时：4ms(intersect)/4ms(intersect2)
内存消耗：3.5MB(intersect)2.8MB(intersect2)

*/

func intersect2(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	res := make([]int, 0)
	for p1, p2 := 0, 0; p1 < len(nums1) && p2 < len(nums2); {
		if nums1[p1] == nums2[p2] {
			res = append(res, nums1[p1])
			p1++
			p2++
		} else if nums1[p1] > nums2[p2] {
			p2++
		} else {
			p1++
		}
	}
	return res
}

func intersect(nums1 []int, nums2 []int) []int {
	m1, m2 := make(map[int]int), make(map[int]int)
	for i := range nums1 {
		m1[nums1[i]]++
	}
	for i := range nums2 {
		m2[nums2[i]]++
	}
	if len(m2) < len(m1) {
		m1, m2 = m2, m1
	}
	res := make([]int, 0)
	for k := range m1 {
		if _, ok := m2[k]; ok {
			for i := 0; i < min(m1[k], m2[k]); i++ {
				res = append(res, k)
			}
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
