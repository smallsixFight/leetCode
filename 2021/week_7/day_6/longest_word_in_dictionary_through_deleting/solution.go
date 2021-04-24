package longest_word_in_dictionary_through_deleting

import "sort"

/*
524
来源：[leetCode](https://leetcode-cn.com/)
题目：[通过删除字母匹配到字典里最长单词](https://leetcode-cn.com/problems/longest-word-in-dictionary-through-deleting/)
标签：`排序` `双指针`

思路：大致处理逻辑整理如下：

- 遍历字符串 `s`，记录每个字母以及对应的索引；
- 对字符串字典 `d` 进行排序；
- 初始化一个空字符串 `res` 作为结果，遍历字符串字典 `d`，当长度比 `res` 长时，尝试在 `s` 查找是否存在子字符串，存在则更新结果 `res`。
- 遍历结束后返回 `res`。

官网运行结果记录
执行用时：36ms/40ms
内存消耗：7.7MB

*/

func findLongestWord(s string, dictionary []string) string {
	record := [26][]int{}
	for i := range s {
		record[s[i]-'a'] = append(record[s[i]-'a'], i)
	}
	sort.Strings(dictionary)
	res := -1
	for i := range dictionary {
		if res == -1 || len(dictionary[i]) > len(dictionary[res]) {
			for _, j := range record[dictionary[i][0]-'a'] {
				p1, p2 := j, 0
				for p1 < len(s) && p2 < len(dictionary[i]) {
					if len(s)-p1 < len(dictionary[i])-p2 {
						break
					}
					if s[p1] == dictionary[i][p2] {
						p2++
					}
					p1++
				}
				if p2 == len(dictionary[i]) {
					res = i
					break
				}
			}
		}
	}
	if res == -1 {
		return ""
	}
	return dictionary[res]
}
