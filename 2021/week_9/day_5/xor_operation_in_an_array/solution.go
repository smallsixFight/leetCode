package same_tree

/*
1486 每日一题
来源：[leetCode](https://leetcode-cn.com/)
题目：[数组异或操作](https://leetcode-cn.com/problems/xor-operation-in-an-array/)
标签：`位运算` `数组`

简单思路：简单地遍历进行异或运算。

官网运行结果记录
执行用时：0ms
内存消耗：1.9MB

*/

func xorOperation(n int, start int) int {
	res := start
	for i := 1; i < n; i++ {
		res = res ^ (start + 2*i)
	}
	return res
}
