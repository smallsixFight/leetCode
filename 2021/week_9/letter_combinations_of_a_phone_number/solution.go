package letter_combinations_of_a_phone_number

/*
17
来源：[leetCode](https://leetcode-cn.com/)
题目：[电话号码的字母组合](https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/)
标签：`深度优先探索` `递归` `字符串` `回溯算法`

简单思路：先使用一个二维数组记录不同数组对应的所有字母，将 `digits` 中第一个数字对应的字母放入到结果集 res 中，然后从索引 1 位置遍历字符串 `digits`，组合所有的可能结果，放入 res 中。遍历结束后返回 res。

这道题个人倾向与 BFS 来处理，而不是 DFS。

官网运行结果记录
执行用时：0ms
内存消耗：2MB

*/

var dict = []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	res := make([]string, 0)
	for i := range dict[digits[0]-'2'] {
		res = append(res, string(dict[digits[0]-'2'][i]))
	}
	for i := 1; i < len(digits); i++ {
		l := len(res)
		for j := 0; j < l; j++ {
			for k := range dict[digits[i]-'2'] {
				res = append(res, res[j]+string(dict[digits[i]-'2'][k]))
			}
		}
		res = res[l:]
	}
	return res
}
