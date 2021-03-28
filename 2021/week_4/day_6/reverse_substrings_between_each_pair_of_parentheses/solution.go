package reverse_substrings_between_each_pair_of_parentheses

/*
1190
来源：[leetCode](https://leetcode-cn.com/)
题目：[反转每对括号间的子串](https://leetcode-cn.com/problems/reverse-substrings-between-each-pair-of-parentheses/)
标签：`栈`

简单思路：从头到右遍历，遇到左括号则在栈 s 中添加一个新的空字符串；遇到英文字符，把连续的字符组成字符串放入栈 s 中；遇到右括号把栈顶的字符串 pop 反转后跟栈顶的字符串合并为一个字符串。遍历结束后，将栈 s 内的字符串组合成一个新字符串返回。

官网运行结果记录
执行用时：0ms
内存消耗：2MB

*/

func reverseParentheses(s string) string {
	stack := [][]byte{{}}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, []byte{})
			continue
		} else if s[i] == ')' {
			bs := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			for j := len(bs) - 1; j >= 0; j-- {
				stack[len(stack)-1] = append(stack[len(stack)-1], bs[j])
			}
			continue
		}
		for i < len(s) && 'a' <= s[i] && s[i] <= 'z' {
			stack[len(stack)-1] = append(stack[len(stack)-1], s[i])
			i++
		}
		i--
	}
	res := make([]byte, 0, len(s))
	for i := range stack {
		res = append(res, stack[i]...)
	}
	return string(res)
}
