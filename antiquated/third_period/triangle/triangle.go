package triangle

func MinimumTotal(triangle [][]int) int {
	if triangle == nil || triangle[0] == nil {
		return 0
	}
	dp := make([]int, len(triangle))
	dp[0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		for j := len(triangle[i]) - 1; j >= 0; j-- {
			if j == len(triangle[i])-1 {
				dp[j] = dp[j-1] + triangle[i][j]
			} else {
				if j == 0 || dp[j-1] > dp[j] {
					dp[j] += triangle[i][j]
				} else {
					dp[j] = dp[j-1] + triangle[i][j]
				}
			}
		}
	}
	minSum := dp[0]
	for i := range dp {
		if dp[i] < minSum {
			minSum = dp[i]
		}
	}
	return minSum
}
