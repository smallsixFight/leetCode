package leaf_similar_trees

import (
	"strings"
)

/*
12 每日一题
来源：[leetCode](https://leetcode-cn.com/)
题目：[整数转罗马数字](https://leetcode-cn.com/problems/integer-to-roman/)
标签：`数学` `数组`

简单思路：看完题目整理之后，可以发现罗马数字的 `4、9、40、90、400、900` 应该作为独立的字符，这个思路能够转换成功的话就很容易做出这道题了。

先为阿拉伯数字和罗马数字做一个映射，包括前面提到的特别数字也需要，然后遍历 num 找到能够减去的最大数字，取对应的罗马数字放入结果集中。

官网运行结果记录
执行用时：8ms
内存消耗：3.3MB

*/

func intToRoman(num int) string {
	nums := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	romans := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	res := strings.Builder{}
	for i := len(nums) - 1; i >= 0; i-- {
		for num >= nums[i] {
			num -= nums[i]
			res.WriteString(romans[i])
		}
	}
	return res.String()
}
