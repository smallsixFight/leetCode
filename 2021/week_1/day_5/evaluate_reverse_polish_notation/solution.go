package evaluate_reverse_polish_notation

import "strconv"

/*
来源：[leetCode](https://leetcode-cn.com/)
题目：[逆波兰表达式求值](https://leetcode-cn.com/problems/evaluate-reverse-polish-notation/)
标签：`栈`

简单思路：使用栈来保存数字。遍历遇到数字则入栈，遇到运算符号则取出栈顶两个数字进行运算后结果入栈，最后返回结果。

官网运行结果记录
执行用时：4ms
内存消耗：3.9MB

*/

func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens))
	for i := range tokens {
		if tokens[i] == "*" || tokens[i] == "/" || tokens[i] == "+" || tokens[i] == "-" {
			v1 := stack[len(stack)-1]
			v2 := 0
			if len(stack) > 1 {
				v2 = stack[len(stack)-2]
				stack = stack[:len(stack)-2]
			} else {
				stack = stack[:len(stack)-1]
			}
			switch tokens[i] {
			case "*":
				stack = append(stack, v1*v2)
			case "/":
				stack = append(stack, v2/v1)
			case "+":
				stack = append(stack, v2+v1)
			case "-":
				stack = append(stack, v2-v1)
			}
		} else {
			v, _ := strconv.ParseInt(tokens[i], 10, 32)
			stack = append(stack, int(v))
		}
	}
	return stack[0]
}
