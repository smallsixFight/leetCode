package tag_validator

import (
	"strings"
)

/*
来源：[leetCode](https://leetcode-cn.com/)
题目：[标签验证器](https://leetcode-cn.com/problems/tag-validator/)
标签：`栈` `字符串`

简单思路：根据题目的详细描述，可以使用栈来处理，从左向右遍历：

- 遇到 `<` 向后遍历获取标签名称，当标签名存在不是大写字母亦或者长度大于 9 或小于 1，返回 false，标签名正确则入栈；
- 遇到 `</` 向后获取标签名称，一样不符合前一点的规则的话就返回 false，否则判断是否与栈顶标签名称相同的，不同则直接返回 false，否则栈顶元素出栈，然后继续向后遍历；
- 遇到 `<![CDATA[` 标签时，先查找后续是否有 `]]>` 标签，没有则直接返回 false，因为标签名已经不符合要求。
- 多个标签如果没有用一个父标签包含起来，也返回 false；
- 遍历完成后，栈如果不为空，返回 false。

官网运行结果记录
执行用时：0ms
内存消耗：2MB

*/

func isValid(code string) bool {
	stack := make([]string, 0)
	isPop := false
	if len(code) == 0 || code[0] != '<' {
		return false
	}
	temp := strings.Builder{}
	for i := 0; i < len(code); i++ {
		if code[i] == '<' {
			if i+1 == len(code) { // 括号右侧至少要有标签名字，否则不符合条件
				return false
			}
			if code[i+1] == '!' {
				if i+11 >= len(code) || code[i:i+9] != "<![CDATA[" { // CDATA 标签判断，若不是 CDATA 标签，感叹号不符合条件，直接返回
					return false
				}
				j := i + 9
				for j < len(code) { // CDATA 结束标签判断，没有直接返回 false
					if j+3 >= len(code) {
						return false
					}
					if code[j:j+3] == "]]>" {
						break
					}
					j++
				}
				if len(stack) == 0 { // 没有父标签进行包括
					isPop = true
				}
				i = j + 2
			} else if code[i+1] == '/' || ('A' <= code[i+1] && code[i+1] <= 'Z') {
				j := i + 1
				if code[j] == '/' {
					j++
				}
				for j < len(code) {
					if 'A' <= code[j] && code[j] <= 'Z' {
						temp.WriteByte(code[j])
						if temp.Len() > 9 { // 长度不符合条件
							return false
						}
						j++
					} else if code[j] == '>' {
						if temp.Len() == 0 { // 长度不符合条件
							return false
						}
						if code[i+1] == '/' {
							// 没有起始标签与当前的结束标签匹配，或者起始标签与结束标签不匹配
							if len(stack) == 0 || stack[len(stack)-1] != temp.String() {
								return false
							}
							stack = stack[:len(stack)-1]
							if len(stack) == 0 { // 没有父标签进行包括
								isPop = true
							}
						} else {
							if isPop { // 没有父标签进行包括，返回 false
								return false
							}
							stack = append(stack, temp.String())
						}
						temp.Reset()
						break
					} else {
						return false
					}
				}
				i = j
			} else { // 不是 CDATA 标签或大写字母，直接返回 false
				return false
			}
		} else if len(stack) == 0 { // 没有标签包裹，但出现 content，返回 false
			return false
		}
	}
	return len(stack) == 0 && temp.Len() == 0
}
