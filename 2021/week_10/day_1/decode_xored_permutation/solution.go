package leaf_similar_trees

/*
1734 每日一题
来源：[leetCode](https://leetcode-cn.com/)
题目：[解码异或后的排列](https://leetcode-cn.com/problems/decode-xored-permutation/)
标签：`位运算`

简单思路：这道题有两个关键的条件，一个是 `n` 为奇数，另外一个是 `perm` 是前 n 个正整数的排列。
那么可以先计算前 n 个数字异或的结果值和 `perm[1]^...perm[n]` 的结果值，这样就能通过异或计算出 perm[0]，然后就像之前做过的异或计算题目一样，依次计算出 perm[i] 即可。

官网运行结果记录
执行用时：184ms(76%)/192ms(26%)
内存消耗：9.8MB(84%)

*/

func decode(encoded []int) []int {
	total, odd := 0, 0
	for i := 1; i <= len(encoded)+1; i++ {
		total ^= i
	}
	for i := 1; i < len(encoded); i += 2 {
		odd ^= encoded[i]
	}
	perm := make([]int, len(encoded)+1)
	perm[0] = total ^ odd
	for i := 0; i < len(encoded); i++ {
		perm[i+1] = encoded[i] ^ perm[i]
	}
	return perm
}
