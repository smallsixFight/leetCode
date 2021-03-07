package binary_tree_preorder_traversal

import "testing"

func Test_preorderTraversal(t *testing.T) {
	n3 := &TreeNode{Val: 3}
	n2 := &TreeNode{Val: 2, Left: n3}
	n1 := &TreeNode{Val: 1, Right: n2}
	t.Log(preorderTraversal(n1))
}

func Test_preorderTraversal2(t *testing.T) {
	n3 := &TreeNode{Val: 3}
	n2 := &TreeNode{Val: 2, Left: n3}
	n1 := &TreeNode{Val: 1, Right: n2}
	t.Log(preorderTraversal2(n1))
}
