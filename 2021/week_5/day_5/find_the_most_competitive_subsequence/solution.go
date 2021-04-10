package find_the_most_competitive_subsequence

/*
1673
来源：[leetCode](https://leetcode-cn.com/)
题目：[找出最具竞争力的子序列](https://leetcode-cn.com/problems/find-the-most-competitive-subsequence/)
标签：`栈` `堆` `贪心算法` `队列`

思路：这道题可以使用一个另类的单调栈来做，这个单调栈近似于一个单调递增栈，因为新的元素除了要比栈顶元素小，还需要后续的长度加上栈内元素的长度需要满足 k 的要求。大致处理逻辑如下：

- 使用一个栈 s 保存结果集，从左向右遍历 nums；
- 若栈为空或 nums[i] 大于等于栈顶元素且栈的大小 < k，则 nums[i] 入栈，否则跳过；
- 若 nums[i] 比栈顶元素小，判断 nums[i] 开始后续的元素数量能否把栈填充到 k，可以则移除栈顶元素，重复前面的步骤；直到栈为空，或不能达到 k 值，或大于等于栈顶元素，将 nums[i] 入栈。
- 遍历结束后，返回栈 s 内元素生成的字符串。

官网运行结果记录
执行用时：208ms(63%)/200ms(97%)
内存消耗：9.4MB(51%)/8.8MB(95%)

*/

func mostCompetitive(nums []int, k int) []int {
	stack := make([]int, 0)
	for i := range nums {
		if len(stack) == 0 || nums[i] >= stack[len(stack)-1] && len(stack) < k {
			stack = append(stack, nums[i])
		} else if nums[i] < stack[len(stack)-1] {
			unUse := len(nums) - i
			for len(stack) > 0 && nums[i] < stack[len(stack)-1] && len(stack)+unUse > k {
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, nums[i])
		}
	}
	return stack
}
