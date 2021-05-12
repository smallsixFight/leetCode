package leaf_similar_trees

import (
	"github.com/leetCode/2021/common"
)

/*
113
来源：[leetCode](https://leetcode-cn.com/)
题目：[路径总和 II](https://leetcode-cn.com/problems/path-sum-ii/)
标签：`树` `深度优先探索`

简单思路：这道题其实也都可以用 BFS 和 DFS 来做，DFS 最好直接用递归。

官网运行结果记录
执行用时：4ms
内存消耗：4.4MB

*/

var res [][]int

func pathSum(root *common.TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}
	arr := make([]int, 0)
	res = make([][]int, 0)
	dfs(root, targetSum, arr)
	return res
}

func dfs(node *common.TreeNode, targetSum int, arr []int) {
	// 要考虑负数
	if node == nil || (node.Val != targetSum && node.Left == nil && node.Right == nil) {
		return
	}
	arr = append(arr, node.Val)
	if node.Val == targetSum && node.Left == nil && node.Right == nil {
		tempArr := make([]int, len(arr))
		copy(tempArr, arr)
		res = append(res, tempArr)
		return
	}
	dfs(node.Left, targetSum-node.Val, arr)
	dfs(node.Right, targetSum-node.Val, arr)
}
