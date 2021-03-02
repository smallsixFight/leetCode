package permutations_ii

func Permutation(nums []int) [][]int {
	res := make([][]int, 0)
	res = append(res, []int{nums[0]})
	for i := 1; i < len(nums); i++ {
		for j := 0; j < len(res); j++ {
			res[j] = append(res[j], nums[i])
		}
		resLen := len(res)
		for j := 0; j < resLen; j++ {
			for x := 0; x < len(res[j])-1; x++ {
				if res[j][x] == res[j][len(res[j])-1] {
					break
				}
				temp := make([]int, len(res[j]))
				copy(temp, res[j])
				temp[x], temp[len(temp)-1] = temp[len(temp)-1], temp[x]
				res = append(res, temp)
			}
		}
	}
	return res
}
