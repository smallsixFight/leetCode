package binary_tree_inorder_traversal

import "testing"

func Test_inorderTraversal(t *testing.T) {
	n3 := &TreeNode{Val: 3}
	n2 := &TreeNode{Val: 2, Left: n3}
	root := TreeNode{Val: 1, Right: n2}
	t.Log(inorderTraversal(&root))
}

func Test_inorderTraversal2(t *testing.T) {
	n3 := &TreeNode{Val: 3}
	n2 := &TreeNode{Val: 2, Left: n3}
	root := TreeNode{Val: 1, Right: n2}
	t.Log(inorderTraversal2(&root))
}
