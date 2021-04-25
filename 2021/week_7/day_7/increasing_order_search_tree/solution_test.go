package longest_word_in_dictionary_through_deleting

import "testing"

func Test_increasingBST(t *testing.T) {
	n1, n2, n3, n4 := &TreeNode{Val: 1}, &TreeNode{Val: 2}, &TreeNode{Val: 3}, &TreeNode{Val: 4}
	n5, n6, n7, n8, n9 := &TreeNode{Val: 5}, &TreeNode{Val: 6}, &TreeNode{Val: 7}, &TreeNode{Val: 8}, &TreeNode{Val: 9}
	n2.Left, n3.Left, n3.Right, n5.Left, n5.Right = n1, n2, n4, n3, n6
	n6.Right, n8.Left, n8.Right = n8, n7, n9
	res := increasingBST(n5)
	for res != nil {
		t.Log(res.Val)
		res = res.Right
	}
}
