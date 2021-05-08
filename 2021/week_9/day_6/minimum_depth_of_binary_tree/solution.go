package same_tree

import (
	"github.com/leetCode/2021/common"
)

/*
111
来源：[leetCode](https://leetcode-cn.com/)
题目：[二叉树的最小深度](https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/)
标签： `树` `深度优先探索` `广度优先探索`

简单思路：这道用 BFS 来处理，当遇到第一个子结点都为 nil 时，直接返回当前遍历到的层数（即高度）。

官网运行结果记录
执行用时：244ms(71%)
内存消耗：18.4MB(74%)

*/

func minDepth(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	queue, res := []*common.TreeNode{root}, 0
	for len(queue) > 0 {
		l := len(queue)
		res++
		for i := 0; i < l; i++ {
			n := queue[i]
			if n.Left == nil && n.Right == nil {
				return res
			}
			if n.Left != nil {
				queue = append(queue, n.Left)
			}
			if n.Right != nil {
				queue = append(queue, n.Right)
			}
		}
		queue = queue[l:]
	}
	return res
}
