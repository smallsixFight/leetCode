package subsets

var res [][]int

func Subsets(nums []int) [][]int {
	res = make([][]int, 1)
	for i := 0; i < len(nums); i++ {
		res = append(res, []int{nums[i]})
	}
	for i := 1; i < len(nums); i++ {
		dfs(res[i], i, nums)
	}
	return res
}

func dfs(arr []int, s int, nums []int) {
	if s == len(nums) {
		return
	}
	for i := s; i < len(nums); i++ {
		temp := make([]int, len(arr))
		copy(temp, arr)
		temp = append(temp, nums[i])
		res = append(res, temp)
		dfs(temp, i+1, nums)
	}
}
