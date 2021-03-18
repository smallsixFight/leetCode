package daily_temperatures

/*
来源：[leetCode](https://leetcode-cn.com/)
题目：[每日温度](https://leetcode-cn.com/problems/daily-temperatures/)
标签：`栈` `哈希表`

简单思路：这道题可以用单调递减栈来做。从左向右遍历：

- 如果栈为空，则入栈；
- 如果小于等于栈顶元素，则入栈；
- 如果大于栈顶元素，则记录栈顶元素的要等待的天数并出栈，然后继续与新的栈顶元素比较，直到遇到大于等于的栈顶元素或栈为空，然后把当前元素入栈；
- 遍历结束后，栈内的元素出栈并设置值为 0。

官网运行结果记录
执行用时：64ms/68ms/72ms/76ms
内存消耗：7.3MB

*/

func dailyTemperatures(T []int) []int {
	stack := make([]int, 0)
	for i := range T {
		for len(stack) > 0 && T[i] > T[stack[len(stack)-1]] {
			T[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	for i := range stack {
		T[stack[i]] = 0
	}
	return T
}
