package subsets_ii

import "sort"

var res [][]int

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res = make([][]int, 1)
	pr := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		res = append(res, []int{nums[i]})
		pr = append(pr, i+1)
	}
	for k := range pr {
		dfs(res[k+1], pr[k], nums)
	}
	return res
}

func dfs(arr []int, s int, nums []int) {
	if s == len(nums) {
		return
	}
	for i := s; i < len(nums); i++ {
		if i != s && nums[i] == nums[i-1] {
			continue
		}
		temp := make([]int, len(arr))
		copy(temp, arr)
		temp = append(temp, nums[i])
		res = append(res, temp)
		dfs(temp, i+1, nums)
	}
}
