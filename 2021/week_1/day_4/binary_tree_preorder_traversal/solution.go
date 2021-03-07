package binary_tree_preorder_traversal

/*
来源：[leetCode](https://leetcode-cn.com/)
题目：[二叉树的前序遍历](https://leetcode-cn.com/problems/binary-tree-preorder-traversal/)
标签：`栈` `树`

简单思路：前序遍历即先处理当前结点，然后处理左子树，最后再处理右子树。因为当前结点已经被处理，不用回溯处理
当前结点，所以迭代的逻辑比中序遍历简单些。因为当前结点左结点的子结点的比右结点的优先级高，所以要先将右结点入栈。

官网运行结果记录
执行用时：0ms
内存消耗：2MB

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, p.Val)
		if p.Right != nil {
			stack = append(stack, p.Right)
		}
		if p.Left != nil {
			stack = append(stack, p.Left)
		}
	}
	return res
}

func preorderTraversal2(root *TreeNode) (vals []int) {
	stack := []*TreeNode{}
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			vals = append(vals, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return
}
