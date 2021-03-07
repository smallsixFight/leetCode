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

func Test_inorderTraversal3(t *testing.T) {
	n6 := &TreeNode{Val: 6}
	n5 := &TreeNode{Val: 5}
	n4 := &TreeNode{Val: 4}
	n3 := &TreeNode{Val: 3, Left: n6}
	n2 := &TreeNode{Val: 2, Left: n4, Right: n5}
	root := TreeNode{Val: 1, Left: n2, Right: n3}
	t.Log(inorderTraversal3(&root))
}
