package binary_tree_right_side_view

import "testing"

func TestRightSideView3(t *testing.T) {
	n4 := &TreeNode{Val: 4}
	n3 := &TreeNode{Val: 3}
	n2 := &TreeNode{Val: 2, Left: n4}
	n1 := &TreeNode{Val: 1, Left: n2, Right: n3}
	t.Log(RightSideView(n1))
}

func TestRightSideView2(t *testing.T) {
	n4 := &TreeNode{Val: 4}
	n3 := &TreeNode{Val: 3, Left: n4}
	n2 := &TreeNode{Val: 2}
	n1 := &TreeNode{Val: 1, Left: n2, Right: n3}
	t.Log(RightSideView(n1))
}

func TestRightSideView(t *testing.T) {
	n5 := &TreeNode{Val: 5}
	n4 := &TreeNode{Val: 4}
	n2 := &TreeNode{Val: 2, Right: n5}
	n3 := &TreeNode{Val: 3, Right: n4}
	n1 := &TreeNode{Val: 1, Left: n2, Right: n3}
	t.Log(RightSideView(n1))
}
