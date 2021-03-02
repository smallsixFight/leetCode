package minimum_depth_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BFS
func MinDepthTwo(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	res := 0
	stack := make([]*TreeNode, 1)
	stack[0] = root
	l := len(stack)
	for l > 0 {
		res++
		for i := 0; i < l; i++ {
			node := stack[0]
			stack = stack[1:]
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			if node.Left == nil && node.Right == nil {
				return res
			}
		}
		l = len(stack)
	}
	return res
}

// DFS
func MinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return dfs(root, 0, 1)
}

func dfs(node *TreeNode, high, minHigh int) int {
	high++
	if node.Left == nil && node.Right == nil {
		if minHigh == 1 || (high > 1 && high < minHigh) {
			minHigh = high
		}
		return minHigh
	}
	if node.Left != nil {
		minHigh = dfs(node.Left, high, minHigh)
	}
	if node.Right != nil {
		minHigh = dfs(node.Right, high, minHigh)
	}
	return minHigh
}
