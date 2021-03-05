package binary_tree_postorder_traversal

import "testing"

func Test_postorderTraversal2(t *testing.T) {
	n3 := &TreeNode{Val: 3}
	n2 := &TreeNode{Val: 2, Left: n3}
	n1 := &TreeNode{Val: 1, Right: n2}
	t.Log(postorderTraversal2(n1)) // 3 2 1
}

func Test_postorderTraversal(t *testing.T) {
	n3 := &TreeNode{Val: 3}
	n2 := &TreeNode{Val: 2, Left: n3}
	n1 := &TreeNode{Val: 1, Right: n2}
	t.Log(postorderTraversal(n1)) // 3 2 1
}
