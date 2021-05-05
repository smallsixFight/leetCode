package letter_combinations_of_a_phone_number

import "testing"

func Test_isValidBST(t *testing.T) {
	n3, n4, n5, n6, n7 := &TreeNode{Val: 3}, &TreeNode{Val: 4}, &TreeNode{Val: 5}, &TreeNode{Val: 6}, &TreeNode{Val: 7}
	n5.Left, n5.Right = n4, n6
	n6.Left, n6.Right = n3, n7
	t.Log(isValidBST(n5)) // false
}

func Test_isValidBST2(t *testing.T) {
	n3, n4, n5, n6, n7 := &TreeNode{Val: 3}, &TreeNode{Val: 4}, &TreeNode{Val: 5}, &TreeNode{Val: 6}, &TreeNode{Val: 7}
	n5.Left, n5.Right = n4, n6
	n6.Left, n6.Right = n3, n7
	t.Log(isValidBST2(n5)) // false
}
