package sort_list

/*
242
来源：[leetCode](https://leetcode-cn.com/)
题目：[有效的字母异位词](https://leetcode-cn.com/problems/valid-anagram/)
标签：`排序` `哈希表`

简单思路：如果 `s` 和 `t` 长度不等，直接返回 false。先遍历 `s`，将字符保存到一个 hashTable 里，并记录出现的次数，然后遍历 `t`，若存在不属于 hashTable 的字符或次数为 0，直接返回 false。遍历结束后直接返回 true。

若包含 unicode 字符的话，通过内循环获取完整的 unicode 字符保存即可（其他非 unicode 字符的可能也需要转换成 unicode 字符）。

官网运行结果记录
执行用时：12ms(isAnagram)/0ms(isAnagram2)
内存消耗：2.8MB(isAnagram)/2.8MB(isAnagram2)

*/

func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	record := [26]int{}
	for i := range s {
		record[s[i]-'a']++
	}
	for i := range t {
		if record[t[i]-'a'] == 0 {
			return false
		}
		record[t[i]-'a']--
	}
	return true
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[byte]int)
	for i := range s {
		m[s[i]]++
	}
	for i := range t {
		if v, ok := m[t[i]]; !ok || v == 0 {
			return false
		}
		m[t[i]]--
	}
	return true
}
