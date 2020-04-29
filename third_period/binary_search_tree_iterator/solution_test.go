package binary_search_tree_iterator

import "testing"

func TestConstructor(t *testing.T) {
	n9 := &TreeNode{Val: 9}
	n20 := &TreeNode{Val: 20}
	n3 := &TreeNode{Val: 3}
	n15 := &TreeNode{Val: 15, Left: n9, Right: n20}
	n7 := &TreeNode{Val: 7, Left: n3, Right: n15}
	c := Constructor(n7)
	for c.HasNext() {
		t.Log(c.Next())
	}
}
