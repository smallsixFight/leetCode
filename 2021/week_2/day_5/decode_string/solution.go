package decode_string

/*
来源：[leetCode](https://leetcode-cn.com/)
题目：[字符串解码](https://leetcode-cn.com/problems/decode-string/)
标签：`栈` `深度优先探索`

简单思路：使用两个栈，s1 用于保存重复的次数，s2 保存连续的子字符串，处理逻辑大概如下：

- 遍历开始前，s2 先新增一个空字符串；
- 遇到数字时，向 s1 push 完整的数字，s2 新增一个空的字符串，然后跳过左括号；
- 遇到字母时，内循环读取连续的字母字符串，push 到 s2，s 向前回溯一个位置；
- 遇到右括号时，取出 s1 栈顶数字 num，再取出 s2 栈顶的字符串，重复 num 次加入 s2 新的栈顶字符串后面。
- 遍历结束后，s2 中的字符串逐个出栈，然后将除了栈底的字符串按出栈的顺序逆序加入栈底字符串的后面，最后返回栈底的字符串。

思路重新整理：
一开始的思路有些错误，在遇到 `3[z]2[2[y]pq4[2[jk]e1[f]]]ef` 这个测试案例后才重新调整思路，就是一对括号里面的字符应该组合在一起才行，内嵌的括号也一样，应该以整合括号内的内容作为整理，而不是一个连续的字符串，
所以，在遇到新的数字时，应开辟一个独立的空间将后面括号内的字符串处理完成，再添加到上层的字符串中。

官网运行结果记录
执行用时：0ms
内存消耗：2MB

*/

func decodeString(s string) string {
	if len(s) == 0 {
		return ""
	}
	s1, s2 := make([]int, 0), [][]byte{{}}
	for i := 0; i < len(s); i++ {
		if '0' <= s[i] && s[i] <= '9' {
			temp := int(s[i] - '0')
			j := i + 1
			for ; j < len(s) && '0' <= s[j] && s[j] <= '9'; j++ {
				temp = temp*10 + int(s[j]-'0')
			}
			s1 = append(s1, temp)
			s2 = append(s2, []byte{})
			i = j // 跳过左括号，不用回溯
			continue
		}
		if 'a' <= s[i] && s[i] <= 'z' {
			temp := []byte{s[i]}
			j := i + 1
			for ; j < len(s) && 'a' <= s[j] && s[j] <= 'z'; j++ {
				temp = append(temp, s[j])
			}
			s2[len(s2)-1] = append(s2[len(s2)-1], temp...)
			i = j - 1 // 需要回溯，在下一个循环中处理
			continue
		}
		if s[i] == ']' {
			num, bs := s1[len(s1)-1], s2[len(s2)-1]
			s1, s2 = s1[:len(s1)-1], s2[:len(s2)-1]
			for j := 0; j < num; j++ {
				s2[len(s2)-1] = append(s2[len(s2)-1], bs...)
			}
		}
	}
	for i := 1; i < len(s2); i++ {
		s2[0] = append(s2[0], s2[i]...)
	}
	return string(s2[0])
}
