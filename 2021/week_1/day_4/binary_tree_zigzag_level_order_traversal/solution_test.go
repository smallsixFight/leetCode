package binary_tree_zigzag_level_order_traversal

import "testing"

func Test_zigzagLevelOrder(t *testing.T) {
	root1 := &TreeNode{Val: 3}
	n1 := TreeNode{Val: 9}
	n2 := TreeNode{Val: 20}
	n3 := TreeNode{Val: 15}
	n4 := TreeNode{Val: 7}
	root1.Left = &n1
	root1.Right = &n2
	n2.Left = &n3
	n2.Right = &n4
	t.Log(zigzagLevelOrder(root1))
}
