package exclusive_time_of_functions

/*
来源：[leetCode](https://leetcode-cn.com/)
题目：[函数的独占时间](https://leetcode-cn.com/problems/exclusive-time-of-functions/)
标签：`栈`

简单思路：这道题的示例说明描述有些问题，例如第一条日志的说明不应该用 `结束` 这个词，容易让人误解，不过确实也是一种解题思路。
这道题可以用栈来模拟实现，从左向右遍历 logs：
- 当 logs[i] 包含 `start` 时，栈顶（不为空）线程的独占时间暂停，更新栈顶线程的独占时间 `logs[i].timestamp - prev.timestamp`，然后将 logs[i] 入栈；
- 当 logs[i] 包含 `end` 时，表示栈顶线程运行结束，出栈并用 `logs[i].timestamp - prev.timestamp + 1` 累计独占时间。

官网运行结果记录
执行用时：4ms/8ms
内存消耗：4.9MB

*/

func exclusiveTime(n int, logs []string) []int {
	prev := 0
	stack := make([]int, 0)
	res := make([]int, n)
	getLogInfo := func(i int) (idx int, status byte, t int) {
		j := 1
		idx = int(logs[i][0] - '0')
		for ; logs[i][j] != ':'; j++ {
			idx = idx*10 + int(logs[i][j]-'0')
		}
		status = logs[i][j+1]
		nn := 1
		for k := len(logs[i]) - 1; logs[i][k] != ':'; k-- {
			t += int(logs[i][k]-'0') * nn
			nn *= 10
		}
		if i == 0 {
			prev = t
		}
		return
	}
	for i := range logs {
		idx, status, t := getLogInfo(i)
		if status == 's' {
			if len(stack) != 0 {
				res[stack[len(stack)-1]] += t - prev
			}
			stack = append(stack, idx)
			prev = t
		} else {
			res[stack[len(stack)-1]] += t - prev + 1 // 包含当前时间所以需要 +1
			stack = stack[:len(stack)-1]
			prev = t + 1 // 当前时间已被独占，所以需要 +1
		}
	}
	return res
}
