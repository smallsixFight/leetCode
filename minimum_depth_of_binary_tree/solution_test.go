package minimum_depth_of_binary_tree

import "testing"

func TestMinDepthTwo(t *testing.T) {
	n15 := &TreeNode{Val: 15}
	n7 := &TreeNode{Val: 7}
	n20 := &TreeNode{Val: 20, Left: n15, Right: n7}
	n9 := &TreeNode{Val: 9}
	n3 := &TreeNode{Val: 3, Left: n9, Right: n20}
	t.Log(MinDepthTwo(n3))
}

func TestMinDepth(t *testing.T) {
	n15 := &TreeNode{Val: 15}
	n7 := &TreeNode{Val: 7}
	n20 := &TreeNode{Val: 20, Left: n15, Right: n7}
	n9 := &TreeNode{Val: 9}
	n3 := &TreeNode{Val: 3, Left: n9, Right: n20}
	t.Log(MinDepth(n3))
}

func TestMinDepth2(t *testing.T) {
	n2 := &TreeNode{Val: 2}
	n1 := &TreeNode{Val: 1, Left: n2}
	t.Log(MinDepth(n1))
}

func TestMinDepth3(t *testing.T) {
	n2 := &TreeNode{Val: 2}
	n1 := &TreeNode{Val: 1, Right: n2}
	t.Log(MinDepth(n1))
}

func TestMinDepth4(t *testing.T) {
	n5 := &TreeNode{Val: 5}
	n4 := &TreeNode{Val: 4}
	n3 := &TreeNode{Val: 3, Left: n5}
	n2 := &TreeNode{Val: 2, Left: n4}
	n1 := &TreeNode{Val: 1, Left: n2, Right: n3}
	t.Log(MinDepth(n1))
}
