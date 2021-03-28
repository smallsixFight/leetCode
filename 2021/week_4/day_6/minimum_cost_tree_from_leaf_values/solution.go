package minimum_cost_tree_from_leaf_values

import "math"

/*
1130
来源：[leetCode](https://leetcode-cn.com/)
题目：[叶值的最小代价生成树](https://leetcode-cn.com/problems/minimum-cost-tree-from-leaf-values/)
标签：`栈` `树` `动态规划`

思路：这道题没做出来，树的结构给我造成了复杂的想象，题目标签的 DP 让我想到先从下往上构建树，但并没能找到如何去组合结点来生成树。
然后去看了题解 [为什么单调递减栈的算法可行](https://leetcode-cn.com/problems/minimum-cost-tree-from-leaf-values/solution/wei-shi-yao-dan-diao-di-jian-zhan-de-suan-fa-ke-xi/)，理解后做了出来。
最后，回归到 DP 的方向，可以发现，其实单调栈的思路也可用来作为 DP 实现的基础，只是写法有些不同。

官网运行结果记录
执行用时：0ms
内存消耗：2MB

*/

func mctFromLeafValues(arr []int) int {
	res := 0
	stack := []int{math.MaxInt32}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := 0; i < len(arr); i++ {
		for len(stack) > 0 && arr[i] > stack[len(stack)-1] {
			res += stack[len(stack)-1] * min(stack[len(stack)-2], arr[i])
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, arr[i])
	}
	for len(stack) > 2 {
		res += stack[len(stack)-1] * stack[len(stack)-2]
		stack = stack[:len(stack)-1]
	}
	return res
}
