package construct_binary_tree_from_inorder_and_postorder_traversal

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

func BuildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	var idx int
	node := &TreeNode{Val: postorder[len(postorder)-1]}
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == postorder[len(postorder)-1] {
			idx = i
			break
		}
	}
	node.Left = BuildTree(inorder[:idx], postorder[:idx])
	node.Right = BuildTree(inorder[idx+1:], postorder[idx:len(postorder)-1])
	return node
}
