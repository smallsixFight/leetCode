package balanced_binary_tree

import "testing"

func TestIsBalanced(t *testing.T) {
	n15 := &TreeNode{Val: 15}
	n7 := &TreeNode{Val: 7}
	n20 := &TreeNode{Val: 20, Left: n15, Right: n7}
	n9 := &TreeNode{Val: 9}
	n3 := &TreeNode{Val: 3, Left: n9, Right: n20}
	t.Log(IsBalanced(n3))
}

func TestIsBalanced2(t *testing.T) {
	n41 := &TreeNode{Val: 4}
	n42 := &TreeNode{Val: 4}
	n31 := &TreeNode{Val: 3, Left: n41, Right: n42}
	n32 := &TreeNode{Val: 3}
	n21 := &TreeNode{Val: 2, Left: n31, Right: n32}
	n22 := &TreeNode{Val: 2}
	n1 := &TreeNode{Val: 1, Left: n21, Right: n22}
	t.Log(IsBalanced(n1))
}
