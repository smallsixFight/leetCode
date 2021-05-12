package xor_queries_of_a_subarray

/*
1310 每日一题
来源：[leetCode](https://leetcode-cn.com/)
题目：[子数组异或查询](https://leetcode-cn.com/problems/xor-queries-of-a-subarray/)
标签：`位运算`

简单思路：用暴力解法的时候还在想这道题怎么就算中等题了，提交后超过时间限制才意识到需要处理。

接着，就想到获取前 n 项的前缀和，由于异或运算，可以通过 `record[0, m] ^ record[0, n] = record[m, n]`，不要做暴力解法的那些重复的运算。

官网运行结果记录
执行用时：68ms(87%)
内存消耗：8.7MB(79%)

*/

func xorQueries(arr []int, queries [][]int) []int {
	record, total := make([]int, len(arr)), 0
	for i := 0; i < len(arr); i++ {
		total ^= arr[i]
		record[i] = total
	}
	res := make([]int, len(queries))
	for i := range queries {
		if queries[i][0] == 0 {
			res[i] = record[queries[i][1]]
		} else {
			res[i] = record[queries[i][1]] ^ record[queries[i][0]-1]
		}

	}
	return res
}
