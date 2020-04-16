package path_sum_ii

import "testing"

func TestPathSum(t *testing.T) {
	n7 := &TreeNode{Val: 7}
	n2 := &TreeNode{Val: 2}
	n51 := &TreeNode{Val: 5}
	n1 := &TreeNode{Val: 1}
	n11 := &TreeNode{Val: 11, Left: n7, Right: n2}
	n41 := &TreeNode{Val: 4, Left: n51, Right: n1}
	n42 := &TreeNode{Val: 4, Left: n11}
	n13 := &TreeNode{Val: 13}
	n8 := &TreeNode{Val: 8, Left: n13, Right: n41}
	n52 := &TreeNode{Val: 5, Left: n42, Right: n8}
	t.Log(PathSum(n52, 22))
}
