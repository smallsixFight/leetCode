package validate_stack_sequences

/*
946
来源：[leetCode](https://leetcode-cn.com/)
题目：[验证栈序列](https://leetcode-cn.com/problems/validate-stack-sequences/)
标签：`栈`

简单思路：使用一个栈 s 和 一个出栈指针 p 来实现，s 保存 pushed 的元素，p 指向当前出栈的元素，处理逻辑如下：

- 从左向右遍历 pushed，如果遇到的元素与 p 指向的元素不同，则放入 s；
- 如果遇到的元素与 p 指向的元素相同，p 向后移动一位，继续与栈顶元素匹配，相同则 s 进行 pop，p 继续后移，直到 s 为空或栈顶元素与 p 指向的元素不同；
- 遍历结束后，如果 s 不会空，则继续与 p 指向的元素比较出栈，如果出现不同的，返回 false。

官网运行结果记录
执行用时：8ms
内存消耗：3.8MB

*/

func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) == 0 {
		return true
	}
	s, p := make([]int, 0), 0
	for i := range pushed {
		if pushed[i] == popped[p] {
			p++
			for len(s) > 0 && s[len(s)-1] == popped[p] {
				s = s[:len(s)-1]
				p++
			}
		} else {
			s = append(s, pushed[i])
		}
	}
	for len(s) > 0 {
		if s[len(s)-1] != popped[p] {
			return false
		}
		p++
		s = s[:len(s)-1]
	}
	return true
}
