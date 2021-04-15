package boolean_evaluation_lcci

/*
面试题 08.14
来源：[leetCode](https://leetcode-cn.com/)
题目：[布尔运算](https://leetcode-cn.com/problems/boolean-evaluation-lcci/)
标签：`栈` `字符串`

思路：注释，后面重新整理。

官网运行结果记录
执行用时：0ms
内存消耗：2.1MB

*/

func countEval(s string, result int) int {
	dp := make([][][2]int, len(s)/2+1)
	for i := 0; i < len(dp); i++ { // 先把对应位置的数字数量登记起来
		dp[i] = make([][2]int, len(s)/2+1)
		dp[i][i][int(s[i*2]-'0')] = 1
	}
	for l := 2; l <= len(dp); l++ { // 从长度 2 开始累计
		for i := 0; i+l-1 < len(dp); i++ { // 分别累计不同数字作为起点，在指定长度各种计算结果的累计数量
			j := i + l - 1
			for k := i; k < j; k++ { // 累计
				r := s[2*k+1]
				switch r {
				case '&':
					dp[i][j][0] += dp[i][k][0]*dp[k+1][j][0] + dp[i][k][0]*dp[k+1][j][1] + dp[i][k][1]*dp[k+1][j][0]
					dp[i][j][1] += dp[i][k][1] * dp[k+1][j][1]
				case '|':
					dp[i][j][0] += dp[i][k][0] * dp[k+1][j][0]
					dp[i][j][1] += dp[i][k][1]*dp[k+1][j][1] + dp[i][k][0]*dp[k+1][j][1] + dp[i][k][1]*dp[k+1][j][0]
				default:
					dp[i][j][0] += dp[i][k][1]*dp[k+1][j][1] + dp[i][k][0]*dp[k+1][j][0]
					dp[i][j][1] += dp[i][k][0]*dp[k+1][j][1] + dp[i][k][1]*dp[k+1][j][0]
				}
			}
		}
	}
	return dp[0][len(dp)-1][result]
}
