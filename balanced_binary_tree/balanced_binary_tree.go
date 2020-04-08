package balanced_binary_tree

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	isB, _ := dfs(root)
	return isB
}

func dfs(node *TreeNode) (bool, int) {
	if node == nil {
		return true, 0
	}
	lh, rh := 1, 1
	isLB, h1 := dfs(node.Left)
	lh += h1
	if !isLB {
		return false, 0
	}
	isRB, h2 := dfs(node.Right)
	rh += h2
	if !isRB {
		return false, 0
	}
	if math.Abs(float64(lh-rh)) > 1 {
		return false, 0
	} else {
		return true, maxVal(lh, rh)
	}
}

func maxVal(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 错误
func dfs2(node *TreeNode) bool {
	if node.Left == nil {
		if node.Right == nil {
			return true
		}
		if (node.Right.Left != nil && (node.Right.Left.Left != nil || node.Right.Left.Right != nil)) ||
			(node.Right.Right != nil && (node.Right.Right.Left != nil || node.Right.Right.Right != nil)) {
			return false
		}
	} else if node.Right == nil {

		if (node.Left.Left != nil && (node.Left.Left.Left != nil || node.Left.Left.Right != nil)) ||
			(node.Left.Right != nil && (node.Left.Right.Left != nil || node.Left.Right.Right != nil)) {
			return false
		}
	} else {
		if dfs2(node.Left) {
			return dfs2(node.Right)
		}
	}
	return false
}
