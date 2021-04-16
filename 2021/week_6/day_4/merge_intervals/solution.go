package merge_intervals

import "sort"

/*
56
来源：[leetCode](https://leetcode-cn.com/)
题目：[合并区间](https://leetcode-cn.com/problems/merge-intervals/)
标签：`排序` `数组`

简单思路：先根据 `intervals` 的第一个元素进行排序，相同的话再按照第二个元素进行排序。接着遍历判断是否可合并，不可合并的放入结果集，可合并的修改结果集最后一个数字的范围。

官网运行结果记录
执行用时：8ms/12ms/16ms
内存消耗：4.5MB

*/

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})
	res := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > res[len(res)-1][1] {
			res = append(res, intervals[i])
		} else if res[len(res)-1][1] < intervals[i][1] {
			res[len(res)-1][1] = intervals[i][1]
		}
	}
	return res
}
