package check_if_word_is_valid_after_substitutions

/*
1003
来源：[leetCode](https://leetcode-cn.com/)
题目：[检查替换后的词是否有效](https://leetcode-cn.com/problems/check-if-word-is-valid-after-substitutions/)
标签：`栈` `字符串`

简单思路：从左向右遍历，当遇到 `a`、`b` 时，入栈，当遇到 `c` 时，判断栈顶两个字符能否组成 `ab`，可以则出栈，继续遍历，不能则直接返回 false。
遍历结束后，栈为空返回 true， 否则返回 false。

官网运行结果记录
执行用时：0ms/4ms
内存消耗：3.6MB

*/

func isValid(s string) bool {
	stack := make([]byte, 0, len(s)/2)
	for i := range s {
		if s[i] == 'c' {
			if len(stack) > 1 && string(stack[len(stack)-2:]) == "ab" {
				stack = stack[:len(stack)-2]
			} else {
				return false
			}
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
