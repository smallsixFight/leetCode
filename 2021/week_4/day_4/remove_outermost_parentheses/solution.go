package remove_outermost_parentheses

/*
1021
来源：[leetCode](https://leetcode-cn.com/)
题目：[删除最外层的括号](https://leetcode-cn.com/problems/remove-outermost-parentheses/)
标签：`栈`

简单思路：从左向右遍历，当遇到左括号，且前一个括号也为左括号时，开始保存字符，并记录左括号个数，使用内循环遍历直到左括号刚好被消耗为 0，那么退出内循环，进行下一次解析，直到遍历结束。

官网运行结果记录
执行用时：0ms
内存消耗：2.4MB

*/

func removeOuterParentheses(S string) string {
	res := make([]byte, 0, len(S))
	for i := 1; i < len(S); i++ {
		if S[i] == '(' && S[i-1] == '(' {
			count, j := 2, i+1
			res = append(res, S[i])
			for ; j < len(S) && count > 0; j++ {
				if S[j] == '(' {
					count++
				} else {
					count--
				}
				if count > 0 {
					res = append(res, S[j])
				}
			}
			i = j - 1
		}
	}
	return string(res)
}
