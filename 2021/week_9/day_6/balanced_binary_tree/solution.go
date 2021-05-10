package same_tree

import (
	"github.com/leetCode/2021/common"
)

/*
110
来源：[leetCode](https://leetcode-cn.com/)
题目：[平衡二叉树](https://leetcode-cn.com/problems/balanced-binary-tree/)
标签：`深度优先探索` `树` `递归`

简单思路：深度遍历分别获取左子树和右子树的高度，比较高度差。

官网运行结果记录
执行用时：8ms
内存消耗：5.7MB

*/

func isBalanced(root *common.TreeNode) bool {
	if root == nil {
		return true
	}
	return dfs(root) >= 0
}

func dfs(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	left, right := dfs(root.Left), dfs(root.Right)
	if left == -1 || right == -1 { // 左右子树其中一个不是平衡二叉树，意味改树也不是平衡二叉树
		return -1
	}
	if abs(left-right) > 1 {
		return -1
	}
	return 1 + max(left, right)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
