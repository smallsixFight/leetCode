package sort_list

import "sort"

/*
274
来源：[leetCode](https://leetcode-cn.com/)
题目：[H 指数](https://leetcode-cn.com/problems/h-index/)
标签：`排序` `哈希表`

简单思路：先对 `citations` 进行排序，接着要找到最大的 `h` 值，从头向尾遍历，可以知道 `h = citations.length - i`，只要找到第一个 `citations.length - i <= citations[i]` 的 `h` 值就行。

这道题真考验语文能力。

官网运行结果记录
执行用时：0ms
内存消耗：2.2MB

*/

func hIndex(citations []int) int {
	sort.Ints(citations)
	for i := range citations {
		h := len(citations) - i
		if h <= citations[i] {
			return h
		}
	}
	return 0
}
