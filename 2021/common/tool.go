package common

import "strconv"

func GetTreeStrings(tree *TreeNode) []string {
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
