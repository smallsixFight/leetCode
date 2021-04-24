package longest_word_in_dictionary_through_deleting

import (
	"sort"
)

/*
767
来源：[leetCode](https://leetcode-cn.com/)
题目：[重构字符串](https://leetcode-cn.com/problems/reorganize-string/)
标签：`堆` `贪心算法` `排序` `字符串`

思路：遍历 `S`，记录各个字母的个数，创建一个长度为 S.length 的 []byte，并按照字母出现数量的降序遍历记录（避免可以分隔开的字母相邻），将字母间隔意味填充到 []byte 中，如果填充的时候发现相邻字母相同，直接返回空字符串；遍历结束返回 []byte 生成的字符串。

官网运行结果记录
执行用时：0ms
内存消耗：2MB

*/

func reorganizeString(S string) string {
	record := make([][2]int, 26)
	for i := range S {
		record[S[i]-'a'][0] = int(S[i])
		record[S[i]-'a'][1]++
	}
	sort.Slice(record, func(i, j int) bool {
		return record[i][1] > record[j][1]
	})
	//fmt.Println(record)
	res, idx := make([]byte, len(S)), 0
	for i := range record {
		for j := 0; j < record[i][1]; j++ {
			if idx >= len(res) {
				idx = 1
			}
			res[idx] = byte(record[i][0])
			if (idx > 0 && res[idx-1] == res[idx]) || (idx < len(res)-1 && res[idx+1] == res[idx]) {
				return ""
			}
			idx += 2
		}
	}
	return string(res)
}
