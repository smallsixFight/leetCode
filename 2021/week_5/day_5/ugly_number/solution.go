package submissions

/*
263 （每日一题）
来源：[leetCode](https://leetcode-cn.com/)
题目：[ugly-number](https://leetcode-cn.com/problems/ugly-number/)
标签：`数学`

思路：如果是 2/3/5 的倍数，则作相应的除法，如果除之后结果值为 1，则返回 true，如果不是 2/3/5 的倍数，直接返回 false。


官网运行结果记录
执行用时：4ms
内存消耗：2.1MB

*/

func isUgly(n int) bool {
	for n > 0 {
		if n == 1 {
			return true
		}
		if n%2 == 0 {
			n /= 2
		} else if n%3 == 0 {
			n /= 3
		} else if n%5 == 0 {
			n /= 5
		} else {
			return false
		}
	}
	return false
}
