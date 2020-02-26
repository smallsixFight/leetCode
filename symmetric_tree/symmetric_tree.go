package symmetric_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsSymmetric(root *TreeNode) bool {
	if root != nil {
		return isEqual(root.Left, root.Right)
	}
	return true
}

func isEqual(p, q *TreeNode) bool {
	if p == q {
		return true
	}
	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}
	if p.Val == q.Val {
		return isEqual(p.Left, q.Right) && isEqual(p.Right, q.Left)
	}
	return false
}
