package unique_paths

// 只需要保存当前行和上一行的数据
func UniquePathsTwo(m, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	pre, cur := make([]int, n), make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				cur[j] = 1
			} else {
				cur[j] = pre[j] + cur[j-1]
			}
		}
		pre = cur
	}
	return cur[n-1]
}

func UniquePaths(m, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for i := 1; i < n; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
