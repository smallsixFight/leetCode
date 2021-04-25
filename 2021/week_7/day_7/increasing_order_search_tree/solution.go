package longest_word_in_dictionary_through_deleting

/*
897 每日一题
来源：[leetCode](https://leetcode-cn.com/)
题目：[递增顺序搜索树](https://leetcode-cn.com/problems/increasing-order-search-tree/)
标签：`树` `深度优先探索` `递归`

简单思路：创建一个新结点作为返回值，按照中序遍历的方式，将结点添加到新结点的后面。

官网运行结果记录
执行用时：0ms
内存消耗：2.2MB

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	top := &TreeNode{}
	p := top
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		if node.Left != nil {
			stack = append(stack, node.Left)
			node.Left = nil // 不再重复遍历左结点
		} else {
			p.Right = node
			p = p.Right
			stack = stack[:len(stack)-1]
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
	}
	return top.Right
}
