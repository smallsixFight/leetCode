package merge_intervals

import (
	"sort"
)

func Merge(intervals [][]int) [][]int {
	res := make([][]int, 0)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i := 0; i < len(intervals); {
		val := intervals[i]
		j := i + 1
		for j < len(intervals) && val[1] >= intervals[j][0] {
			if intervals[j][1] > val[1] {
				val[1] = intervals[j][1]
			}
			j++
		}
		i = j
		res = append(res, val)
	}
	return res
}
