package basic_calculator

import (
	"strconv"
	"strings"
)

/*
来源：[leetCode](https://leetcode-cn.com/)
题目：[基本计算器](https://leetcode-cn.com/problems/basic-calculator/)
标签：`栈` `数学`

简单思路：之前做过 `逆波兰表达式求值`，所以想着先将算式转换成逆波兰表达式，再进行求值。转换思路如下：
- 使用一个字符串和一个栈，字符串用来存逆波兰表达式（tokens），一个用来临时存储运算符号（stack）；
- 遇到数字时进行内循环获取完整的数字，然后存入 tokens；
- 遇到运算符号时，如果 stack 不为空并且栈顶不是左括号，则取出栈顶的元素加入 tokens 中，否则 push 到 stack 中；
- 遇到右括号时，则不断从 stack 中 push 运算符加入 token 中，知道遇见左括号。

最后将 s1 进行逆波兰表达式求值。

官网运行结果记录
执行用时：12ms/16ms
内存消耗：8.5MB
todo: 待学习优化

*/

func calculate(s string) int {
	tokens, stack := make([]string, 0, len(s)), make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		if len(tokens) == 0 && s[i] == '+' {
			continue
		}
		if '0' <= s[i] && s[i] <= '9' {
			temp := strings.Builder{}
			temp.WriteByte(s[i])
			j := i + 1
			for ; j < len(s) && s[j] >= '0' && s[j] <= '9'; j++ {
				temp.WriteByte(s[j])
			}
			i = j - 1
			tokens = append(tokens, temp.String())
		} else if s[i] == '+' || s[i] == '-' { // 给的都是有效的表达式，所以不用做额外的判断
			if len(stack) > 0 && stack[len(stack)-1] != '(' {
				v := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				tokens = append(tokens, string(v))
			}
			stack = append(stack, s[i])
		} else if s[i] == '(' {
			stack = append(stack, s[i])
		} else if s[i] == ')' {
			for j := len(stack) - 1; stack[j] != '('; j-- {
				tokens = append(tokens, string(stack[j]))
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
		}
	}
	for i := len(stack) - 1; i >= 0; i-- {
		tokens = append(tokens, string(stack[i]))
	}
	return evalRPN(tokens)
}

func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens))
	for i := range tokens {
		if tokens[i] == "+" || tokens[i] == "-" {
			v1 := stack[len(stack)-1]
			v2 := 0
			if len(stack) > 1 {
				v2 = stack[len(stack)-2]
				stack = stack[:len(stack)-2]
			} else {
				stack = stack[:len(stack)-1]
			}
			switch tokens[i] {
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
