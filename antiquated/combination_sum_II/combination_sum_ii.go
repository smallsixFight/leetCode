package combination_sum_II

import "sort"

var cArr []int
var res [][]int

func Solution(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return nil
	}
	sort.Ints(candidates)
	cArr = candidates
	res = make([][]int, 0)
	for i := 0; i < len(cArr) && cArr[i] <= target; i++ {
		if i > 0 && cArr[i] == cArr[i-1] {
			continue
		}
		dfs(i, target, nil)
	}
	return res
}

func dfs(i, t int, arr []int) {
	if cArr[i] > t {
		return
	}
	if cArr[i] == t {
		temp := make([]int, len(arr))
		copy(temp, arr)
		res = append(res, append(temp, cArr[i]))
		return
	}
	t -= cArr[i]
	arr = append(arr, cArr[i])
	for j := i + 1; j < len(cArr); j++ {
		dfs(j, t, arr)
		for j+1 < len(cArr) && cArr[j+1] == cArr[j] {
			j++
		}
	}
}
