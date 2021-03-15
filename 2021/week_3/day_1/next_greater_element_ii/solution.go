package next_greater_element_ii

/*
来源：[leetCode](https://leetcode-cn.com/)
题目：[下一个更大元素 II](https://leetcode-cn.com/problems/next-greater-element-ii/)
标签：`栈`

简单思路：这道题可以用与 `下一个更大元素 I` 处理 num2 的思路来处理，可以对 nums 进行两次遍历处理，逻辑如下：

- 从左向右遍历，并使用一个 stack 保存遍历的数字；
- 当遍历到的数字大于栈顶的数字时，对栈顶的数字记录并出栈，然后继续与新的栈顶数字比较，直到栈顶数字大于等于遍历到的数字，然后将遍历到的数字入栈；
- 第一次遍历结束后，如果栈不为空（此时栈是单调递减栈），以栈顶数字的位置作为结束位置，再次遍历 nums，这次不需要入栈，只需要再次与栈内数字做比较，若找到大于栈顶的数字，则记录并出栈；
- 二次遍历结束后，栈还不会空，依次出栈，并记录更大的数字为 -1。
- 然后返回结果集。

官网运行结果记录
执行用时：24ms/28ms/32ms（28ms 击败 96%，32ms 击败 83%）
内存消耗：6.6MB（击败 76%）

*/

func nextGreaterElements(nums []int) []int {
	stack, res := make([]int, 0), make([]int, len(nums))
	for i := range nums {
		for len(stack) > 0 && nums[i] > nums[stack[len(stack)-1]] {
			res[stack[len(stack)-1]] = nums[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	if len(stack) > 0 {
		lastIdx := stack[len(stack)-1]
		for i := 0; i < lastIdx; i++ {
			for len(stack) > 0 && nums[i] > nums[stack[len(stack)-1]] {
				res[stack[len(stack)-1]] = nums[i]
				stack = stack[:len(stack)-1]
			}
		}
	}
	for len(stack) > 0 {
		res[stack[len(stack)-1]] = -1
		stack = stack[:len(stack)-1]
	}
	return res
}
