package symmetric_tree

import "testing"

func Test_isSymmetric(t *testing.T) {
	root1 := TreeNode{Val: 1}
	root1.Left = &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}
	root1.Right = &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}}
	t.Log(isSymmetric(&root1)) // true

	root2 := TreeNode{Val: 1}
	root2.Left = &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}
	root2.Right = &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}
	t.Log(isSymmetric(&root2)) // false
}
