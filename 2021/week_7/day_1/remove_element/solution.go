package sort_list

/*
27 每日一题
来源：[leetCode](https://leetcode-cn.com/)
题目：[移除元素](https://leetcode-cn.com/problems/remove-element/)
标签：`数组` `双指针`

简单思路：使用双指针法，处理逻辑如下：

- 初始化两个指针 p1 = -1，p2 = 0；
- 以 p2 作为限制条件，从头向尾遍历 nums；
- 若 `nums[p2] == val`，p2++；
- 若 `nums[p2] != val`，nums[p1+1] = nums[p2]，p1++，p2++；
- 遍历结束后，返回 p1 + 1;

官网运行结果记录
执行用时：0ms
内存消耗：2.1MB

*/

func removeElement(nums []int, val int) int {
	p1, p2 := -1, 0
	for ; p2 < len(nums); p2++ {
		if nums[p2] != val {
			nums[p1+1] = nums[p2]
			p1++
		}
	}
	return p1 + 1
}
