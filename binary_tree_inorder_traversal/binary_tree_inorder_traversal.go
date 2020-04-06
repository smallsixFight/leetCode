package binary_tree_inorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InOrderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	stack := make([]*TreeNode, 0)
	p := root
	for len(stack) > 0 || p != nil {
		isPop := false
		if p == nil {
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			isPop = true
		}
		if !isPop && p.Left != nil {
			stack = append(stack, p)
			p = p.Left
			continue
		}
		res = append(res, p.Val)
		if p.Right != nil {
			p = p.Right
		} else {
			p = nil
		}
	}
	return res
}

func InOrderTraversalWithRecursive(root *TreeNode) []int {
	res := make([]int, 0)
	if root != nil {
		res = append(res, InOrderTraversalWithRecursive(root.Left)...)
		res = append(res, root.Val)
		res = append(res, InOrderTraversalWithRecursive(root.Right)...)
	}
	return res
}
