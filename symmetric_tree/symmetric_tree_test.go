package symmetric_tree

import "testing"

func TestIsSymmetric(t *testing.T) {
	root1 := TreeNode{Val: 1}
	root1.Left = &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}
	root1.Right = &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}}
	t.Log(IsSymmetric(&root1))

	root2 := TreeNode{Val: 1}
	root2.Left = &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}
	root2.Right = &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}
	t.Log(IsSymmetric(&root2))

	root3 := TreeNode{Val: 2}
	root3.Left = &TreeNode{Val: 3, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}
	root3.Right = &TreeNode{Val: 3, Right: &TreeNode{Val: 4}}
	t.Log(IsSymmetric(&root3))
}
