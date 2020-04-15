package construct_binary_tree_from_preorder_and_inorder_traversal

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printTree(tree *TreeNode) []string {
	if tree == nil {
		return nil
	}
	res := make([]string, 0)
	stack := []*TreeNode{tree}
	for len(stack) > 0 {
		l := len(stack)
		for i := 0; i < l; i++ {
			node := stack[i]
			if node != nil {
				res = append(res, strconv.FormatInt(int64(node.Val), 10))
				stack = append(stack, node.Left)
				stack = append(stack, node.Right)
			} else {
				res = append(res, "null")
			}
		}
		stack = stack[l:]
	}
	return res
}

func BuildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) != len(preorder) || len(preorder) == 0 {
		return nil
	}
	var idx int
	node := &TreeNode{Val: preorder[0]}
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			idx = i
			break
		}
	}
	node.Left = BuildTree(preorder[1:idx+1], inorder[:idx])
	node.Right = BuildTree(preorder[idx+1:], inorder[idx+1:])
	return node
}
