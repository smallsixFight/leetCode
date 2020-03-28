package combinations

func Combine(n, k int) [][]int {
	res := make([][]int, 0)
	if n == 0 || k == 0 || n < k {
		return append(res, []int{})
	}
	if n == k {
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			arr[i] = i + 1
		}
		res = append(res, arr)
		return res
	}
	for i := 1; i <= n-k+1; i++ {
		res = append(res, []int{i})
	}
	for len(res[0]) < k {
		l := len(res)
		for i := 0; i < l; i++ {
			val := res[i][len(res[i])-1]
			for j := 1; k-len(res[i]) <= n-val-j+1; j++ { // 剪枝
				temp := make([]int, len(res[i]))
				copy(temp, res[i])
				temp = append(temp, val+j)
				res = append(res, temp)
			}
		}
		res = res[l:]
	}
	return res
}
