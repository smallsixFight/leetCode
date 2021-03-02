package path_sum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// DFS
func HasPathNumTwo(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Val == sum && root.Left == nil && root.Right == nil {
		return true
	}
	return HasPathNumTwo(root.Left, sum-root.Val) || HasPathNumTwo(root.Right, sum-root.Val)
}

// BFS
func HasPathNum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Val == sum && root.Left == nil && root.Right == nil {
		return true
	}
	stack := make([]*TreeNode, 1)
	stack[0] = root
	l := len(stack)
	for l > 0 {
		for i := 0; i < l; i++ {
			node := stack[0]
			stack = stack[1:]
			if node.Left != nil {
				node.Left.Val += node.Val
				if node.Left.Val == sum && node.Left.Left == nil && node.Left.Right == nil {
					return true
				}
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				node.Right.Val += node.Val
				if node.Right.Val == sum && node.Right.Left == nil && node.Right.Right == nil {
					return true
				}
				stack = append(stack, node.Right)
			}
		}
		l = len(stack)
	}
	return false
}
