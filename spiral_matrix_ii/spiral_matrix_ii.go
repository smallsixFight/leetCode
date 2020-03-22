package spiral_matrix_ii

func GenerateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	for k, val := 0, 0; k < (n+1)/2; k++ {
		p := n - k - 1
		for i := k; i <= p; i++ {
			val++
			res[k][i] = val
		}
		for j := k + 1; j < p; j++ {
			val++
			res[j][p] = val
		}
		if p != k {
			for i := p; i >= k; i-- {
				val++
				res[p][i] = val
			}
			for j := p - 1; j > k; j-- {
				val++
				res[j][k] = val
			}
		}
	}
	return res
}
