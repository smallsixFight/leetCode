package symmetric_tree

/*
101
来源：[leetCode](https://leetcode-cn.com/)
题目：[对称二叉树](https://leetcode-cn.com/problems/symmetric-tree/)
标签：`树` `深度优先探索` `广度优先探索`

简单思路：这道题是跟 `some-tree` 这道题很类似，可以稍微调整下 `some-tree` 题解，将左右子结点比较进行对换即可。

官网运行结果记录
执行用时：4ms
内存消耗：2.9MB

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSome(root.Left, root.Right)
}

func isSome(p, q *TreeNode) bool {
	if p == nil && q == p {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	return isSome(p.Left, q.Right) && isSome(p.Right, q.Left)
}
