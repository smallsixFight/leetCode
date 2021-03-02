package flatten_binary_tree_to_linked_list

import "testing"

func TestFlatten(t *testing.T) {
	n3 := &TreeNode{Val: 3}
	n4 := &TreeNode{Val: 4}
	n6 := &TreeNode{Val: 6}
	n2 := &TreeNode{Val: 2, Left: n3, Right: n4}
	n5 := &TreeNode{Val: 5, Right: n6}
	n1 := &TreeNode{Val: 1, Left: n2, Right: n5}
	Flatten(n1)
	for n1 != nil {
		t.Log(n1.Val)
		n1 = n1.Right
	}
}
