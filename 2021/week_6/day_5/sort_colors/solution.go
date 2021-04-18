package sort_colors

/*
75
来源：[leetCode](https://leetcode-cn.com/)
题目：[插入区间](https://leetcode-cn.com/problems/insert-interval/)
标签：`排序` `数组` `双指针`

简单思路：根据题意，需要把 `0` 放到数组开头，`2` 放到数组末尾，大致处理逻辑如下：

- 使用 p1、p2 分别指向 `nums` 开头和尾部；
- 从头向尾遍历 `nums`：
	- 若 `nums[i] == 0 && i != p1`，交换 `nums[i]` 和 `nums[p1]` 且 p1 自增加一；
	- 若 `nums[i] == 2`，交换 `nums[i]` 和 `nums[p2]` 且 p2 减一；
	- 其他情况，i 自增加一。

官网运行结果记录
执行用时：0ms
内存消耗：2MB

*/

func sortColors(nums []int) {
	if len(nums) < 2 {
		return
	}
	p1, p2 := 0, len(nums)-1
	for i := 0; i <= p2 && p1 < p2; {
		if nums[i] == 0 && i != p1 { // 交换后，nums[i] 需要重新比较
			nums[i], nums[p1] = nums[p1], nums[i]
			p1++
		} else if nums[i] == 2 { // 交换后，nums[i] 需要重新比较
			nums[i], nums[p2] = nums[p2], nums[i]
			p2--
		} else {
			i++
		}
	}
}
