package binary_tree_inorder_traversal

import "testing"

func TestInOrderTraversal2(t *testing.T) {
	n10 := &TreeNode{Val: 10}
	n9 := &TreeNode{Val: 9}
	n8 := &TreeNode{Val: 8}
	n7 := &TreeNode{Val: 7}
	n6 := &TreeNode{Val: 6}
	n5 := &TreeNode{Val: 5, Left: n10}
	n4 := &TreeNode{Val: 4, Left: n8, Right: n9}
	n3 := &TreeNode{Val: 3, Left: n6, Right: n7}
	n2 := &TreeNode{Val: 2, Left: n4, Right: n5}
	n1 := &TreeNode{Val: 1, Left: n2, Right: n3}
	// 8 4 9 2 10 5 1 6 3 7
	t.Log(InOrderTraversal(n1))
}

func TestInOrderTraversal(t *testing.T) {
	root := TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}
	t.Log(InOrderTraversal(&root))
}

func TestInOrderTraversalWithRecursive(t *testing.T) {
	root := TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}
	t.Log(InOrderTraversalWithRecursive(&root))
}
