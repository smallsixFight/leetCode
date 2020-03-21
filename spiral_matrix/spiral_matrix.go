package spiral_matrix

func SpiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	res := make([]int, 0)
	m, n := len(matrix), len(matrix[0])
	for k := 0; k < (n+1)/2 && k < (m+1)/2; k++ {
		p1, p2 := n-k-1, m-k-1
		for i := k; i <= p1; i++ { // 取顶行
			res = append(res, matrix[k][i])
		}
		for j := k + 1; j < p2; j++ { // 取右侧列
			res = append(res, matrix[j][p1])
		}
		for i := p1; i >= k && p2 != k; i-- { // 取底行，如果只有一行的情况下要避免重复
			res = append(res, matrix[p2][i])
		}
		for j := p2 - 1; j > k && p1 != k; j-- { // 取左侧列，只有一列的情况下要避免重复
			res = append(res, matrix[j][k])
		}
	}
	return res
}
