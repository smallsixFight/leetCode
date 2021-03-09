package basic_calculator_ii

import "strings"

/*
来源：[leetCode](https://leetcode-cn.com/)
题目：[基本计数器 II](https://leetcode-cn.com/problems/basic-calculator-ii/)
标签：`栈` `设计`

简单思路：使用两个栈，栈一（s1）用来保存数字，栈二（s2）用来保存运算符，并分入栈和计算两种思路，大致逻辑如下：

- 从头遍历，遇到数字 push 到 s1，遇到 `(` push 到 s2;
- 遇到 `*`、`/`，如果 s2 栈顶是 '*'、`/`，先取出栈顶的运算符计算 push 到 s1，再将当前运算符 push 到 s2；
- 遇到 `+`、`-`，s2 逐个出栈，直到遇到 `(`，s2 逐个出栈的同时 s1 出栈计算，最后将新的运算符 push 到 s2；
- 遇到 `)`，s2 逐个出栈，直到遇到 `(`，s2 逐个出栈的同时 s1 出栈计算，最后移除 s2 栈顶的 `(`；
- 遍历结束后，如果 s2 不为空，则逐步出栈计算，最后返回 s1[0] 上的值。

解释：由于遇到 `*` 和 `/` 且 s2 栈顶也是 `*` 或 `/`，就会先将栈顶的取出运算，所以 s2 不会有超过一个的乘号或除号；当遍历到的新运算符是加减号时，它们的优先级是最低的，所以 s2 中的所有运算符号需要先处理完再将遍历到的运算符放入 s2 中。

官网运行结果记录
执行用时：0ms
内存消耗：2.9MB

*/

func calculate(s string) int {
	s = strings.TrimSpace(s)
	if len(s) > 0 && s[0] == '-' {
		s = "-" + s
	}
	s1, s2 := make([]int, 0, len(s)), make([]byte, 0)
	cal := func() {
		v1, v2 := s1[len(s1)-2], s1[len(s1)-1]
		s1 = s1[:len(s1)-1]
		o := s2[len(s2)-1]
		s2 = s2[:len(s2)-1]
		if o == '*' {
			s1[len(s1)-1] = v1 * v2
		} else if o == '/' {
			s1[len(s1)-1] = v1 / v2
		} else if o == '+' {
			s1[len(s1)-1] = v1 + v2
		} else {
			s1[len(s1)-1] = v1 - v2
		}
	}
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		if s[i] == '(' {
			s2 = append(s2, s[i])
			continue
		}
		if '0' <= s[i] && s[i] <= '9' {
			temp := int(s[i] - '0')
			j := i + 1
			for j < len(s) && s[j] >= '0' && s[j] <= '9' {
				temp = temp*10 + int(s[j]-'0')
				j++
			}
			s1 = append(s1, temp)
			i = j - 1
			continue
		}
		if s[i] == '+' || s[i] == '-' {
			for len(s2) > 0 && (s2[len(s2)-1] != '(') {
				cal()
			}
			s2 = append(s2, s[i])
			continue
		}
		if s[i] == '*' || s[i] == '/' {
			if len(s2) > 0 && (s2[len(s2)-1] == '*' || s2[len(s2)-1] == '/') {
				cal()
			}
			s2 = append(s2, s[i])
			continue
		}
		if s[i] == ')' {
			for len(s2) > 0 && s2[len(s2)-1] != '(' {
				cal()
			}
			if len(s2) > 0 {
				s2 = s2[:len(s2)-1]
			}
		}
	}
	for len(s2) > 0 {
		cal()
	}
	return s1[0]
}
