package maximum_depth_of_binary_tree

/*
104
来源：[leetCode](https://leetcode-cn.com/)
题目：[二叉树的最大深度](https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/)
标签：`树` `深度优先探索` `递归`

简单思路：这道题用递归可以很简单地实现 DFS，所以也写了 BFS 的方法。

官网运行结果记录
执行用时：4ms
内存消耗：4.2MB

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BFS
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	res := 0
	for len(queue) > 0 {
		res++
		l := len(queue)
		for i := 0; i < l; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[l:]
	}
	return res
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
