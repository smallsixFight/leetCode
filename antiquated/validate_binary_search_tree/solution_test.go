package validate_binary_search_tree

import "testing"

func TestIsValidBST(t *testing.T) {
	n1 := &TreeNode{Val: 1}
	n3 := &TreeNode{Val: 3}
	root := &TreeNode{Val: 2, Left: n1, Right: n3}
	t.Log(IsValidBST(root))
}

func TestIsValidBST2(t *testing.T) {
	n3 := &TreeNode{Val: 3}
	n6 := &TreeNode{Val: 6}
	n1 := &TreeNode{Val: 1}
	n4 := &TreeNode{Val: 4, Left: n3, Right: n6}

	root := &TreeNode{Val: 5, Left: n1, Right: n4}
	t.Log(IsValidBST(root))
}
