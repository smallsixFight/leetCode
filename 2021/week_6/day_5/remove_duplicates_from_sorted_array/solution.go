package sort_colors

/*
26 每日一题
来源：[leetCode](https://leetcode-cn.com/)
题目：[删除有序数组中的重复项](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/)
标签：`数组` `双指针`

简单思路：根据题意，`nums` 是已经排好序的，那么 `nums` 重复的数字是会连续出现的，那么可以使用双指针法，大致逻辑如下：

- 使用 p1、p2 指向 nums[0]，res 作为结果，然后开始遍历；
- 若 `nums[p1] == nums[p2]`，p2++；
- 若 `nums[p1] != nums[p2]`，nums[p1++] = nums[p2]，p2++，l++；
- 遍历结束后返回 res。


官网运行结果记录
执行用时：8ms
内存消耗：4.5MB

*/

func removeDuplicates(nums []int) int {
	p1, p2, res := 0, 0, 0
	for p2 < len(nums) {
		if nums[p1] != nums[p2] {
			p1++
			nums[p1] = nums[p2]
			res++
		}
		p2++
	}
	return res + 1
}
