package html_entity_parser

import "strings"

/*
1410
来源：[leetCode](https://leetcode-cn.com/)
题目：[HTML 实体解析器](https://leetcode-cn.com/problems/html-entity-parser/submissions/)
标签：`栈` `字符串`

思路：这道题是之前参加一次竞赛做的题，这里重新做一下，以及优化下代码。

不知道 `entityParser` 为什么在最后一个测试案例死活不能过，而且只有提交的时候会出现。

找到问题，是提交用的测试案例有问题。

官网运行结果记录
执行用时：32ms(95%-entityParser2)
内存消耗：6.9MB(50%-entityParser2)

*/

func entityParser2(text string) string {
	sb := strings.Builder{}
	for i := 0; i < len(text); {
		if text[i] == '&' {
			if i+5 < len(text) && text[i:i+6] == "&quot;" {
				sb.WriteByte('"')
				i += 6
			} else if i+5 < len(text) && text[i:i+6] == "&apos;" {
				sb.WriteByte('\'')
				i += 6
			} else if i+4 < len(text) && text[i:i+5] == "&amp;" {
				sb.WriteByte('&')
				i += 5

			} else if i+3 < len(text) && text[i:i+4] == "&gt;" {
				sb.WriteByte('>')
				i += 4
			} else if i+3 < len(text) && text[i:i+4] == "&lt;" {
				sb.WriteByte('<')
				i += 4
			} else if i+6 < len(text) && text[i:i+7] == "&frasl;" {
				sb.WriteByte('/')
				i += 7
			} else {
				sb.WriteByte(text[i])
				i++
			}
		} else {
			sb.WriteByte(text[i])
			i++
		}
	}
	return sb.String()
}

func entityParser(text string) string {
	res := strings.Builder{}
	m := map[string]byte{"&quot;": '"', "&apos;": '\'', "&amp;": '&', "&gt;": '>',
		"&lt;": '<', "&frasl;": '/'}
	for i := 0; i < len(text); i++ {
		if text[i] == '&' {
			j := i + 1
			for text[j] != ';' {
				j++
			}
			if v, ok := m[text[i:j+1]]; ok {
				res.WriteByte(v)
			} else {
				res.WriteString(text[i : j+1])
			}
			i = j
		} else {
			res.WriteByte(text[i])
		}
	}
	return res.String()
}
