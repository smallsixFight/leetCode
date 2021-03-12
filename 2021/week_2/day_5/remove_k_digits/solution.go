package remove_k_digits

/*
来源：[leetCode](https://leetcode-cn.com/)
题目：[移掉 K 位数字](https://leetcode-cn.com/problems/remove-k-digits/)
标签：`栈` `贪心算法`

简单思路：首先，尝试思考如何移除一位使剩下的数字达到最小，尝试后就会发现，只要找到第一个 num[i] > num[i+1]，移除 num[i] 就可以了。然后将这个思路套用在移除 k 个数上即可。

官网运行结果记录
执行用时：0ms
内存消耗：2.6MB

*/

func removeKdigits(num string, k int) string {
	stack := make([]byte, 0)
	for i := range num {
		v := num[i]
		for k > 0 && len(stack) > 0 && v < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, v)
	}
	stack = stack[:len(stack)-k]
	i := 0
	for i < len(stack) && stack[i] == '0' {
		i++
	}
	stack = stack[i:]
	if len(stack) == 0 {
		return "0"
	}
	return string(stack)
}
