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

1+2+3 -(1+2-(1+3))
思路二：由于题目说明给的运算式只有加减运算，那么只需要从左到右进行运算基本上就可以了，只需要考虑括号里面的要优先计算，大致逻辑如下：
- 使用两个 stack 分别存储数字（s1）和存储运算符和括号（s2）；
- 从头遍历，如果当前符号是数字，则获取完整的数字入栈到 s1；
- 如果是左括号，则直接入栈到 s2；
- 如果是 `+` 或 `-`，如果 s2 为空或 s2 栈顶为左括号，则入栈到 s2，否则取出 s1 栈顶两个元素和 s2 的栈顶运算符进行计算，将计算的结果入栈到 s1，将新的运算符放入 s2；
- 如果是右括号，当 s2 不为空，则再获取 s2 栈顶的运算符，计算一个新的值入栈到 s1，栈顶为左括号则直接 pop 掉，不用计算。

关于右括号判断为什么不需要用 for，这是由于前面遇到运算符都会计算，所以 s2 中左括号后面只会有 0 或 1 个运算符，不会有多个。

官网运行结果记录
执行用时：12ms/16ms		4ms（思路二）
内存消耗：8.5MB			3MB（思路二）


*/

func calculate2(s string) int {
	s = strings.TrimSpace(s)
	if len(s) > 0 && s[0] == '-' { // 处理 `-2+ 1` 的特殊情况
		s = "0" + s
	}
	s1, s2 := make([]int, 0, len(s)), make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		if s[i] == '(' {
			s2 = append(s2, s[i])
			continue
		}
		if s[i] >= '0' && s[i] <= '9' {
			num, n := int(s[i]-'0'), 10
			j := i + 1
			for j < len(s) && s[j] >= '0' && s[j] <= '9' {
				num = num*n + int(s[j]-'0')
				j++
			}
			i = j - 1
			s1 = append(s1, num)
			continue
		}
		if s[i] == '+' || s[i] == '-' {
			if len(s2) == 0 || s2[len(s2)-1] == '(' {
				s2 = append(s2, s[i])
			} else {
				op := s2[len(s2)-1]
				v1, v2 := s1[len(s1)-2], s1[len(s1)-1]
				s1 = s1[:len(s1)-2]
				if op == '+' {
					s1 = append(s1, v1+v2)
				} else if op == '-' {
					s1 = append(s1, v1-v2)
				}
				s2[len(s2)-1] = s[i]
			}
			continue
		}
		if s[i] == ')' {
			if len(s2) > 1 && s2[len(s2)-1] != '(' {
				op := s2[len(s2)-1]
				s2 = s2[:len(s2)-1]
				v1, v2 := s1[len(s1)-2], s1[len(s1)-1]
				s1 = s1[:len(s1)-2]
				if op == '+' {
					s1 = append(s1, v1+v2)
				} else if op == '-' {
					s1 = append(s1, v1-v2)
				}
			}
			s2 = s2[:len(s2)-1]
		}
	}
	for len(s2) > 0 {
		op := s2[len(s2)-1]
		s2 = s2[:len(s2)-1]
		v1, v2 := s1[len(s1)-2], s1[len(s1)-1]
		s1 = s1[:len(s1)-2]
		if op == '+' {
			s1 = append(s1, v1+v2)
		} else if op == '-' {
			s1 = append(s1, v1-v2)
		}
	}
	return s1[0]
}

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
