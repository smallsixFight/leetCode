package leaf_similar_trees

import "github.com/leetCode/2021/common"

/*
872
来源：[leetCode](https://leetcode-cn.com/)
题目：[叶子相似的树](https://leetcode-cn.com/problems/leaf-similar-trees/)
标签：`树` `深度优先探索`

简单思路：使用 DFS 分别获取两棵树的叶子结点，然后比较叶子结点。

官网运行结果记录
执行用时：0ms
内存消耗：2.5MB

*/

func leafSimilar(root1 *common.TreeNode, root2 *common.TreeNode) bool {
	s1, s2 := dfs(root1), dfs(root2)
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i].Val != s2[i].Val {
			return false
		}
	}
	return true
}

func dfs(root *common.TreeNode) []*common.TreeNode {
	stack := []*common.TreeNode{root}
	res := make([]*common.TreeNode, 0)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Left == nil && node.Right == nil {
			res = append(res, node)
			continue
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return res
}
