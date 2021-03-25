package remove_all_adjacent_duplicates_in_string

/*
1047
来源：[leetCode](https://leetcode-cn.com/)
题目：[删除字符串中的所有相邻重复项](https://leetcode-cn.com/problems/remove-all-adjacent-duplicates-in-string/)
标签：`栈`

简单思路：从左向右遍历，逐个字符与栈顶字符比较，两个相同则把栈顶元素出栈，否则新元素入栈，最后返回栈内结果。

官网运行结果记录
执行用时：4ms
内存消耗：6MB

*/

func removeDuplicates(S string) string {
	stack := make([]byte, 0, len(S))
	for i := range S {
		if len(stack) != 0 && stack[len(stack)-1] == S[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, S[i])
		}
	}
	return string(stack)
}
