package path_sum

import "testing"

func TestHasPathNumTwo(t *testing.T) {
	n7 := &TreeNode{Val: 7}
	n2 := &TreeNode{Val: 2}
	n1 := &TreeNode{Val: 1}
	n11 := &TreeNode{Val: 11, Left: n7, Right: n2}
	n13 := &TreeNode{Val: 13}
	n41 := &TreeNode{Val: 4, Right: n1}
	n42 := &TreeNode{Val: 4, Left: n11}
	n8 := &TreeNode{Val: 8, Left: n13, Right: n41}
	n5 := &TreeNode{Val: 5, Left: n42, Right: n8}
	t.Log(HasPathNumTwo(n5, 22))
}

func TestHasPathNum(t *testing.T) {
	n7 := &TreeNode{Val: 7}
	n2 := &TreeNode{Val: 2}
	n1 := &TreeNode{Val: 1}
	n11 := &TreeNode{Val: 11, Left: n7, Right: n2}
	n13 := &TreeNode{Val: 13}
	n41 := &TreeNode{Val: 4, Right: n1}
	n42 := &TreeNode{Val: 4, Left: n11}
	n8 := &TreeNode{Val: 8, Left: n13, Right: n41}
	n5 := &TreeNode{Val: 5, Left: n42, Right: n8}
	t.Log(HasPathNum(n5, 23))
}

func TestHasPathNum2(t *testing.T) {
	n2 := &TreeNode{Val: 2}
	n1 := &TreeNode{Val: 1, Left: n2}
	t.Log(HasPathNum(n1, 1))
}
