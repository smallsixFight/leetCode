package minimum_insertions_to_balance_a_parentheses_string

/*
1541
来源：[leetCode](https://leetcode-cn.com/)
题目：[平衡括号字符串的最少插入次数](https://leetcode-cn.com/problems/minimum-insertions-to-balance-a-parentheses-string/)
标签：`栈` `字符串`

思路：这道题跟 `使括号有效的最少添加` 不能说一模一样，只能说很像，由一个左括号与一个右括号进行匹配，变成一个左括号与连续两个右括号进行匹配，大致处理逻辑如下：

- 使用 num1 记录未能匹配的成功的左括号个数，使用 res 作为结果；
- 从左向右遍历，遇到左括号时 num1 自增加一；
- 遇到右括号时，并且是连续的两个右括号：
	- 若 num1 >= 0，num1 减一；
	- 若 num1 == 0，res 加一；
- 遇到单个右括号时：
	- 若 num1 >= 0，num1 减一，res 加一；
	- 若 num1 == 0，res 加 二；
- 遍历结束后，返回 `num1 * 2 + res`。

官网运行结果记录
执行用时：8ms(100%)
内存消耗：6.3MB(100%)

*/

func minInsertions(s string) int {
	num1, res := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			num1++
		} else {
			if i+1 < len(s) && s[i+1] == ')' {
				if num1 > 0 {
					num1--
				} else {
					res++
				}
				i++
			} else {
				if num1 > 0 {
					num1--
					res++
				} else {
					res += 2
				}
			}
		}
	}
	return num1*2 + res
}
