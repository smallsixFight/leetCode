package validate_binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	isOk, _, _ := dfs(root)
	return isOk
}

func dfs(node *TreeNode) (bool, int, int) {
	minV, maxV := node.Val, node.Val
	if node.Left != nil && node.Left.Val >= node.Val {
		return false, 0, 0
	}
	if node.Right != nil && node.Right.Val <= node.Val {
		return false, 0, 0
	}
	if node.Left != nil {
		isOk, minVal, maxVal := dfs(node.Left)
		if !isOk || maxVal >= node.Val {
			return false, 0, 0
		}
		minV = minVal
	}
	if node.Right != nil {
		isOk, minVal, maxVal := dfs(node.Right)
		if !isOk || minVal <= node.Val {
			return false, 0, 0
		}
		maxV = maxVal
	}
	return true, minV, maxV
}
