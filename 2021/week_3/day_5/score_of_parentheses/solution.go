package score_of_parentheses

/*
来源：[leetCode](https://leetcode-cn.com/)
题目：[括号的分数](https://leetcode-cn.com/problems/score-of-parentheses/)
标签：`栈` `字符串`

简单思路：从左向右遍历，使用两个栈 s1、s2 分别保存数值和运算符号，处理逻辑如下：

- 如果 S[i] 为 `(` 时，若下一个为 `)`，s1 放入 1；若下一个为 `(`，s1 放入 2，s2 放入 `*`；
- 如果 S[i] 为 `)` 时，按顺序走下面两步进行处理：
	- 若 S[i-1]是 `)`，说明遇到一个需要 `*2` 的操作，那么需要先讲用括号内需要相加的先加起来，然后再 `*2`；
	- 若 S[i+1]是 `(`，说明下一个需要进行相加，s2 加入一个 `+`。
- 最后，可能存在还需要相加的情况，s2 逐个出栈进行计算即可。

官网运行结果记录
执行用时：0ms
内存消耗：1.9MB

*/

func scoreOfParentheses(S string) int {
	s1, s2 := make([]int, 0), make([]byte, 0)
	cal := func() {
		v1, v2 := s1[len(s1)-1], s1[len(s1)-2]
		if s2[len(s2)-1] == '*' {
			s1[len(s1)-2] = v1 * v2
		} else {
			s1[len(s1)-2] = v1 + v2
		}
		s1 = s1[:len(s1)-1]
		s2 = s2[:len(s2)-1]
	}
	for i := range S {
		if S[i] == '(' {
			if S[i+1] == ')' {
				s1 = append(s1, 1)
			} else {
				s1 = append(s1, 2)
				s2 = append(s2, '*')
			}
		} else {
			if S[i-1] == ')' {
				for s2[len(s2)-1] != '*' { // 先把加号的处理完
					cal()
				}
				cal() // 再 *2 一次
			}
			if i+1 < len(S) && S[i+1] == '(' {
				s2 = append(s2, '+')
			}
		}
	}
	for len(s2) > 0 {
		cal()
	}
	return s1[0]
}
