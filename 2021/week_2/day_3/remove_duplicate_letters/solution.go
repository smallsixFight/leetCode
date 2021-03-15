package remove_duplicate_letters

/*
来源：[leetCode](https://leetcode-cn.com/)
题目：[取出重复的字母](https://leetcode-cn.com/problems/remove-duplicate-letters/)
标签：`栈` `贪心算法` `字符串`

简单思路：使用一个双向链表 l1 存结果，使用一个 hashTable 做字典，表示字符是否已存在 l1 中。从头向后遍历，每个字符都与链表的最后字符作比较，如果比链表最后的字符小并且不存在 l1 中，则遍历后续的字符，如果存在与链尾相同的字符，则移除链尾字符，当前字符再与新的链尾字符比较，直到链尾字符不大于当前字符，如果与链尾字符相同，则忽略，否则加入链尾（实现最小字典序以及保证每个字符都出现在结果集）。
为了避免链尾每次都需要去遍历后面的字符寻找是否有相同的字符，可以使用使用一个 hashTable 记录，当链尾字符需要处理并且没有记录时，则遍历后进行记录，如果链尾字符被移除，则对记录进行 -1 操作。

官网运行结果记录
执行用时：0ms/4ms
内存消耗：2.1MB

*/

func removeDuplicateLetters(s string) string {
	l1, dict, countDict := make([]byte, 0), make([]bool, 26), make([]int, 26)
	for i := range s {
		countDict[int(s[i]-'a')]++
	}
	for i := range s {
		if !dict[s[i]-'a'] {
			for len(l1) > 0 && s[i] < l1[len(l1)-1] {
				idx := l1[len(l1)-1] - 'a'
				if countDict[idx] > 1 {
					l1 = l1[:len(l1)-1]
					countDict[idx]--
					dict[idx] = false
					continue
				}
				break
			}
			dict[s[i]-'a'] = true
			l1 = append(l1, s[i])
		} else if countDict[s[i]-'a'] > 1 {
			countDict[s[i]-'a']--
		}
	}
	return string(l1)
}

//简单思路：使用一个 slice 来保的结果值，再使用一个 hashTable 记录字符出现前的最新位置，再使用一个标识记录当前遇到的最小字符，然后开始从左向右遍历，逻辑如下：
//
//- 如果字符还没遇到过，在 slice 对应的位置记录字符；
//- 如果字符已经出现过，如果这个字符大于记录的最小字符，并且上一个出现的位置在最小字符的前面，那么上一个出现的位置替换为 `A`，记录当前字符，否则当前字符对应的位置记录为 `A`；
//- 每次循环最后，然后比较记录最小的字符；
//- 最后遍历 slice，取出不为 `A` 的字符组成结果。
//
//犯的错误：一开始没怎么注意和去理解 `返回结果的字段序最小字典序`，认为每个字符应位于小于它的字符后面即可。

// abcdeaedcb
//func removeDuplicateLetters(s string) string {
//	res := make([]byte, len(s))
//	record := make(map[byte]int)
//	minRuneIdx := 0
//	for i := 0; i < len(s); i ++ {
//		if idx, ok := record[s[i]]; !ok {
//			record[s[i]] = i
//			res[i] = s[i]
//		} else {
//			if s[i] > s[minRuneIdx] && idx < minRuneIdx {
//				res[idx] = 'A'
//				res[i] = s[i]
//				record[s[i]] = i
//			} else {
//				res[i] = 'A'
//			}
//		}
//		if s[i] < s[minRuneIdx] {
//			minRuneIdx = i
//		}
//	}
//	str := strings.Builder{}
//	for i := range res {
//		if res[i] != 'A' {
//			str.WriteByte(res[i])
//		}
//	}
//	return str.String()
//}
