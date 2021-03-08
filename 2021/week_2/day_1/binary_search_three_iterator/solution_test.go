package binary_search_three_iterator

import "testing"

func TestBSTIterator(t *testing.T) {
	n9 := &TreeNode{Val: 9}
	n20 := &TreeNode{Val: 20}
	n3 := &TreeNode{Val: 3}
	n15 := &TreeNode{Val: 15, Left: n9, Right: n20}
	root := &TreeNode{Val: 7, Left: n3, Right: n15}
	iterator := Constructor(root)
	for iterator.HasNext() {
		t.Log(iterator.Next())
	}
}

func TestBSTIterator2(t *testing.T) {
	n9 := &TreeNode{Val: 9}
	n20 := &TreeNode{Val: 20}
	n3 := &TreeNode{Val: 3}
	n15 := &TreeNode{Val: 15, Left: n9, Right: n20}
	root := &TreeNode{Val: 7, Left: n3, Right: n15}
	iterator := Constructor2(root)
	for iterator.HasNext() {
		t.Log(iterator.Next())
	}
}
