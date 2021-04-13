package minimum_distance_between_bst_nodes

import "math"

/*
783 每日一题
来源：[leetCode](https://leetcode-cn.com/)
题目：[二叉搜索树节点最小距离](https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes/)
标签：`栈` `设计`

思路：根据二叉搜索树的性质，每个节点的最小差值一定在取左子树的最大值以及右子树的最小值分别与当前节点作比较好取较小值。
遍历处理每个非叶子节点，比较获取最终结果值。

官网运行结果记录
执行用时：0ms
内存消耗：2.3MB

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := math.MaxInt32
	ns := []*TreeNode{root}
	for len(ns) > 0 {
		n := ns[0]
		ns = ns[1:]
		if n.Left != nil {
			res = min(res, n.Val-maxVal(n.Left))
			ns = append(ns, n.Left)
		}
		if n.Right != nil {
			res = min(res, minVal(n.Right)-n.Val)
			ns = append(ns, n.Right)
		}
	}
	return res
}

func maxVal(n *TreeNode) int {
	for n.Right != nil {
		n = n.Right
	}
	return n.Val
}

func minVal(n *TreeNode) int {
	for n.Left != nil {
		n = n.Left
	}
	return n.Val
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
