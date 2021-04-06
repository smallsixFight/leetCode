package remove_all_adjacent_duplicates_in_string_ii

/*
1209
来源：[leetCode](https://leetcode-cn.com/)
题目：[删除字符串中的所有相邻重复项 II](https://leetcode-cn.com/problems/remove-all-adjacent-duplicates-in-string-ii/)
标签：`栈`

思路：从左向右遍历，将当前字母和截至当前字母的连续个数放入栈 s 中，如何数量达到 k，则进行移除，否则入栈。最后返回栈内字母组成的字符串。

`removeDuplicates` 将 `stack := make([][2]int32, 0)` 改为 `stack := make([][2]int32, 0, len(s))`，内存占用直接从 6.1MB 降低到 3.5MB。
分别记录字符和出现次数的 `f2`，减少了一次遍历，时间降低到了 0ms。

官网运行结果记录
执行用时：0ms(f2)		4ms(removeDuplicates)
内存消耗：3.4MB(f2)	3.5MB(removeDuplicates)

*/

func f2(s string, k int) string {
	stack := make([]byte, 0, len(s))
	countStack := make([]int32, 0, len(s)-1)
	for i := range s {
		if len(stack) == 0 || s[i] != stack[len(stack)-1] {
			stack = append(stack, s[i])
			countStack = append(countStack, 1)
		} else {
			if countStack[len(countStack)-1]+1 == int32(k) {
				stack = stack[:len(stack)-k+1]
				countStack = countStack[:len(countStack)-k+1]
			} else {
				stack = append(stack, s[i])
				countStack = append(countStack, countStack[len(countStack)-1]+1)
			}
		}
	}
	return string(stack)
}

func removeDuplicates(s string, k int) string {
	stack := make([][2]int32, 0, len(s))
	for i := range s {
		if len(stack) == 0 || int32(s[i]-'a') != stack[len(stack)-1][0] {
			stack = append(stack, [2]int32{int32(s[i] - 'a'), 1})
		} else {
			if stack[len(stack)-1][1]+1 == int32(k) {
				stack = stack[:len(stack)-k+1]
			} else {
				stack = append(stack, [2]int32{int32(s[i] - 'a'), stack[len(stack)-1][1] + 1})
			}
		}
	}
	res := make([]byte, len(stack))
	for i := range stack {
		res[i] = byte(stack[i][0] + 'a')
	}
	return string(res)
}
