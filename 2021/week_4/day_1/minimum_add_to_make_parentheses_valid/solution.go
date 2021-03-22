package minimum_add_to_make_parentheses_valid

/*
来源：[leetCode](https://leetcode-cn.com/)
题目：[使括号有效的最少添加](https://leetcode-cn.com/problems/minimum-add-to-make-parentheses-valid/)
标签：`栈` `贪心算法`

简单思路：由于左右括号要按序成对才有效，可以使用一个记录未匹配到的左括号的变量 num1 和一个记录需要补充的左括号数量的变量 num2 来处理，大致处理逻辑如下：

- 从左向右遍历，如果是左括号，则 num1 += 1；
- 如果是右括号，且 num1 > 0，则 num1 -= 1；如果 num1 = 0，num2 += 1；
- 遍历结束后，返回 `num1 + num2`。

官网运行结果记录
执行用时：0ms
内存消耗：2MB

*/

func minAddToMakeValid(S string) int {
	num1, num2 := 0, 0
	for i := range S {
		if S[i] == '(' {
			num1++
		} else {
			if num1 > 0 {
				num1--
			} else {
				num2++
			}
		}
	}
	return num1 + num2
}
