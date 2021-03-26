package smallest_subsequence_of_distinct_characters

/*
1081
来源：[leetCode](https://leetcode-cn.com/)
题目：[不同字符的最小子序列](https://leetcode-cn.com/problems/smallest-subsequence-of-distinct-characters/)
标签：`栈` `贪心算法` `字符串`

简单思路：这道题跟 `去除重复的字母` 是一模一样的，这里再次整理思路复习一下。处理逻辑如下：

- 先进行一次遍历，记录每个字符重复出现的次数；
- 接着进行第二次遍历，使用一个栈 s 作为结果集和一个 hashTable 记录已经在结果集中的字符，从左向右，进行如下处理：
	- 如果当前字符还未在 s 中，并且字典序小于 s 栈顶的字符，并且栈顶的字符在后续遍历的字符中还会出现，则 pop，并设置 hashTable 为true；
	- 继续执行前一个步骤，直到栈为空 or 栈顶元素字典序小于当前字符 or 栈顶字符在后续不再出现;
	- 内循环遍历结束后，将字符入栈，并在 hashTable 记录为 true；
	- 如果字符已存在，则将次数记录减一。

官网运行结果记录
执行用时：0ms
内存消耗：1.9MB

*/

func smallestSubsequence(s string) string {
	countRecord := make([]int, 26)
	for i := range s {
		countRecord[s[i]-'a']++
	}
	stack, dict := make([]byte, 0), make([]bool, 26)
	for i := range s {
		if !dict[s[i]-'a'] {
			for len(stack) > 0 && s[i] < stack[len(stack)-1] {
				idx := stack[len(stack)-1] - 'a'
				if countRecord[idx] > 1 {
					stack = stack[:len(stack)-1]
					countRecord[idx]--
					dict[idx] = false
					continue
				}
				break
			}
			dict[s[i]-'a'] = true
			stack = append(stack, s[i])
		} else {
			countRecord[s[i]-'a']--
		}
	}
	return string(stack)
}
