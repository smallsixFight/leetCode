package same_tree

/*
690
来源：[leetCode](https://leetcode-cn.com/)
题目：[相同的树](https://leetcode-cn.com/problems/same-tree/)
标签：`树` `深度优先探索`

简单思路：这道题可以通过递归来维护一个天然的栈结构，或者自定义一个栈结构来进行迭代进行 DFS 处理。

官网运行结果记录
执行用时：0ms
内存消耗：2.1MB

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && p == q {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
