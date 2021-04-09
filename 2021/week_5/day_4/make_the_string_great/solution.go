package make_the_string_great

/*
1544
来源：[leetCode](https://leetcode-cn.com/)
题目：[整理字符串](https://leetcode-cn.com/problems/make-the-string-great/)
标签：`栈` `字符串`

思路：使用一个栈 s 保存结果集，然后从左向右遍历，如果栈为空或者 s[i] 与栈顶元素不是大小写字母关系，则将 s[i] 放入栈中，否则移除栈顶元素。最后返回栈 s 内元素组成的字符串。

官网运行结果记录
执行用时：0ms
内存消耗：2.1MB

*/

func makeGood(s string) string {
	abs := func(v int) int {
		if v < 0 {
			return -v
		}
		return v
	}
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && abs(int(s[i])-int(stack[len(stack)-1])) == 32 {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
