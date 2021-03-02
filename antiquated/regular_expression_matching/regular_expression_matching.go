package regular_expression_matching

// DP
func IsMatch2(s, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[len(s)][len(p)] = true
	// i 从 len(s) 是因为空串也可能跟匹配串匹配成功
	for i := len(s); i >= 0; i-- {
		for j := len(p) - 1; j >= 0; j-- {
			firstMatch := i < len(s) && (s[i] == p[j] || p[j] == '.')
			if j+1 < len(p) && p[j+1] == '*' {
				dp[i][j] = dp[i][j+2] || (firstMatch && dp[i+1][j])
			} else {
				dp[i][j] = firstMatch && dp[i+1][j+1]
			}
		}
	}
	return dp[0][0]
}

// backtracking
func IsMatch(s, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	firstMatch := len(s) != 0 && (s[0] == p[0] || p[0] == '.')
	if len(p) > 1 && p[1] == '*' {
		return IsMatch(s, p[2:]) || (firstMatch && IsMatch(s[1:], p))
	} else {
		return firstMatch && IsMatch(s[1:], p[1:])
	}
}
