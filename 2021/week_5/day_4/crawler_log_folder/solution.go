package crawler_log_folder

/*
1598
来源：[leetCode](https://leetcode-cn.com/)
题目：[文件夹操作日志搜集器](https://leetcode-cn.com/problems/crawler-log-folder/)
标签：`栈`

思路：使用一个栈 s 作记录，从左向右遍历，处理逻辑如下：

- 遇到 "x/" 时，放入栈 s 中；
- 遇到 "../" 时，取出栈顶元素，栈为空则跳过；
- 遇到 "./" 时，跳过。

其实可不用栈，可以使用一个常量 num 来记录返回结果即可。

官网运行结果记录
执行用时：0ms/4ms
内存消耗：3MB

*/

func minOperations(logs []string) int {
	res := 0
	for i := range logs {
		if logs[i] == "./" {
			continue
		} else if logs[i] == "../" {
			if res > 0 {
				res--
			}
		} else {
			res++
		}
	}
	return res
}
