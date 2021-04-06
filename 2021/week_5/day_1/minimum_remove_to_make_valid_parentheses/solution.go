package minimum_remove_to_make_valid_parentheses

/*
1249
来源：[leetCode](https://leetcode-cn.com/)
题目：[移除无效的括号](https://leetcode-cn.com/problems/minimum-remove-to-make-valid-parentheses/)
标签：`栈` `字符串`

思路：从左向右遍历，将遍历到的字符放入栈 s，并记录左括号的个数。
- 当遍历到左括号或者字母时，直接入栈；
- 当遍历到右括号时，如果左括号个数为 0，不入栈，如果左括号个数大于 0，入栈并且左括号个数减一；
- 遍历结束后，若左括号的个数大于 0，s 中的字符依次出栈放入结果集，如果遇到左括号，则并不放入结果集，直到左括号个数记录变为 0。

官网运行结果记录
执行用时：8ms(86%)
内存消耗：6MB(100%)

*/

func minRemoveToMakeValid(s string) string {
	stack, lc := make([]byte, 0, len(s)), 0
	for i := range s {
		if s[i] == '(' {
			lc++
		} else if s[i] == ')' {
			if lc == 0 {
				continue
			}
			lc--
		}
		stack = append(stack, s[i])
	}
	for i := len(stack) - 1; i >= 0 && lc > 0; i-- {
		if stack[i] == '(' {
			lc--
			stack = append(stack[:i], stack[i+1:]...)
		}
	}
	return string(stack)
}
