package same_tree

/*
1720 每日一题
来源：[leetCode](https://leetcode-cn.com/)
题目：[解码异或后的数组](https://leetcode-cn.com/problems/decode-xored-array/)
标签：`位运算`

简单思路：根据异或的计算方式，进行对 encoded[i-1] 和 arr[i-1] 其实就可以计算出 arr[i] 的值。

官网运行结果记录
执行用时：40ms(44%)
内存消耗：6.8MB(82%)

*/

func decode(encoded []int, first int) []int {
	res := make([]int, len(encoded)+1)
	res[0] = first
	for i := 1; i < len(res); i++ {
		res[i] = encoded[i-1] ^ res[i-1]
	}
	return res
}
