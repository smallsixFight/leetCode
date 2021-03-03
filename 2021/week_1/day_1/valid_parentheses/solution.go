package valid_parentheses

/*
来源：[leetCode](https://leetcode-cn.com/)
题目：[有效的括号](https://leetcode-cn.com/problems/valid-parentheses/)
标签：`栈` `字符串`
简单思路：从头遍历，遇到左括号入栈，遇到右括号 pop 栈顶的左括号比较，匹配则继续遍历，不匹配则 return false。

些许优化：
1. 当栈里括号的数量大于字符串剩余未匹配的长度时，可直接 return false；
2. 当栈里的数量为 0，遇到右括号时，可直接 return false。

官网运行结果记录
结果：通过
执行用时：0ms
内存消耗：2MB

*/

// 用 slice 模拟栈结构
type stack struct {
	data []byte
}

func (s *stack) pop() byte {
	l := len(s.data)
	if l == 0 {
		return 0
	}
	v := s.data[l-1]
	s.data = s.data[:l-1]
	return v
}

func (s *stack) push(v byte) {
	s.data = append(s.data, v)
}

func (s *stack) length() int {
	return len(s.data)
}

func isValid(s string) bool {
	st := stack{}
	for i := range s {
		if st.length() > len(s[i:]) {
			return false
		}
		if s[i] == ')' || s[i] == '}' || s[i] == ']' {
			if st.length() == 0 {
				return false
			}
			v := st.pop()
			if (v == '(' && s[i] != ')') || (v == '{' && s[i] != '}') || (v == '[' && s[i] != ']') {
				return false
			}
		} else {
			st.push(s[i])
		}
	}
	return st.length() == 0
}
