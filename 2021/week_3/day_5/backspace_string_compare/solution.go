package backspace_string_compare

/*
来源：[leetCode](https://leetcode-cn.com/)
题目：[比较含退格的字符串](https://leetcode-cn.com/problems/backspace-string-compare/)
标签：`栈` `双指针`

简单思路：使用两个栈（s1，s2）分别保存 S、T 遍历后剩余的字符，然后比较是否一致。

进阶思路：因为 `#` 只考虑退格，那么从右向左遍历，遍历到的字符如果不能被退格且不相同，那么直接返回 false。
使用 num1、num2 分别记录 S、T 累计的 `#` 的个数，使用 p1、p2 标识 S、T 中当前需要比较的字符。
若 S 或 T 其中一个先遍历完成，那么另外一个就要判断剩余的字符串能否组成一个空字符串，可以就返回 true，否则返回 false。

官网运行结果记录
执行用时：0ms			0ms(进阶)
内存消耗：2MB(44%)	2MB(100%/92%)

*/

func backspaceCompare2(S string, T string) bool {
	num1, num2 := 0, 0
	p1, p2 := len(S)-1, len(T)-1
	for p1 >= 0 && p2 >= 0 {
		if S[p1] != '#' && num1 == 0 && T[p2] != '#' && num2 == 0 {
			if S[p1] != T[p2] {
				return false
			} else {
				p1--
				p2--
			}
			continue
		}
		if S[p1] == '#' {
			num1++
			p1--
		} else if S[p1] != '#' && num1 > 0 {
			p1--
			num1--
		}
		if T[p2] == '#' {
			num2++
			p2--
		} else if T[p2] != '#' && num2 > 0 {
			p2--
			num2--
		}
	}
	if p1 >= 0 {
		for i := p1; i >= 0; i-- {
			if S[i] == '#' {
				num1++
			} else {
				if num1 == 0 {
					return false
				}
				num1--
			}
		}
	} else if p2 >= 0 {
		for i := p2; i >= 0; i-- {
			if T[i] == '#' {
				num2++
			} else {
				if num2 == 0 {
					return false
				}
				num1--
			}
		}
	}
	return true
}

func backspaceCompare(S string, T string) bool {
	s1, s2 := make([]byte, 0), make([]byte, 0)
	for i := 0; i < len(S) || i < len(T); i++ {
		if i < len(S) {
			if S[i] == '#' {
				if len(s1) > 0 {
					s1 = s1[:len(s1)-1]
				}
			} else {
				s1 = append(s1, S[i])
			}
		}
		if i < len(T) {
			if T[i] == '#' {
				if len(s2) > 0 {
					s2 = s2[:len(s2)-1]
				}
			} else {
				s2 = append(s2, T[i])
			}
		}
	}
	return string(s1) == string(s2)
}
