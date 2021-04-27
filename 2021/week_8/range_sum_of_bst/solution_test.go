package range_sum_of_bst

import "testing"

func Test_rangeSumBST(t *testing.T) {
	n3, n5, n7 := &TreeNode{Val: 3}, &TreeNode{Val: 5}, &TreeNode{Val: 7}
	n10, n15, n18 := &TreeNode{Val: 10}, &TreeNode{Val: 15}, &TreeNode{Val: 18}
	n10.Left, n10.Right = n5, n15
	n5.Left, n5.Right, n15.Right = n3, n7, n18
	t.Log(rangeSumBST(n10, 7, 15)) // 32
}

func Test_rangeSumBST2(t *testing.T) {
	n1, n6, n13 := &TreeNode{Val: 1}, &TreeNode{Val: 6}, &TreeNode{Val: 13}
	n3, n5, n7 := &TreeNode{Val: 3}, &TreeNode{Val: 5}, &TreeNode{Val: 7}
	n10, n15, n18 := &TreeNode{Val: 10}, &TreeNode{Val: 15}, &TreeNode{Val: 18}
	n10.Left, n10.Right = n5, n15
	n5.Left, n5.Right, n15.Left, n15.Right = n3, n7, n13, n18
	n3.Left, n7.Left = n1, n6
	t.Log(rangeSumBST(n10, 6, 10)) // 23
}
