package combination_sum

var res [][]int
var t int
var c []int

func CombinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return res
	}
	res = make([][]int, 0)
	t = target
	c = candidates
	for i := 0; i < len(c); i++ {
		deathExplore(0, i, nil)
	}
	return res
}

func deathExplore(sum, i int, arr []int) {
	sum += c[i]
	if sum > t {
		return
	}
	if sum == t {
		temp := make([]int, len(arr))
		copy(temp, arr)
		temp = append(temp, c[i])
		res = append(res, temp)
		return
	}
	arr = append(arr, c[i])
	for j := i; j < len(c); j++ {
		deathExplore(sum, j, arr)
	}
}
