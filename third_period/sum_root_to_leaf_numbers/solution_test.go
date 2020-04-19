package sum_root_to_leaf_numbers

import "testing"

func TestSumNumbers2(t *testing.T) {
	n5 := &TreeNode{Val: 5}
	n1 := &TreeNode{Val: 1}
	n9 := &TreeNode{Val: 9, Left: n5, Right: n1}
	n0 := &TreeNode{Val: 0}
	n4 := &TreeNode{Val: 4, Left: n9, Right: n0}
	t.Log(SumNumbers(n4))
}

func TestSumNumbers(t *testing.T) {
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	n1 := &TreeNode{Val: 1, Left: n2, Right: n3}
	t.Log(SumNumbers(n1))
}
