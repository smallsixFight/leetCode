package longest_valid_parentheses

func LongestValidParentheses(s string) int {
	if len(s) < 2 {
		return 0
	}
	dp := make([]int, len(s))
	maxLen := 0
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			idx := i - dp[i-1] - 1
			if idx >= 0 && s[idx] == '(' {
				dp[i] = dp[i-1] + 2
				if idx > 0 {
					dp[i] += dp[idx-1]
				}
				if dp[i] > maxLen {
					maxLen = dp[i]
				}
			}
		}
	}
	return maxLen
}
