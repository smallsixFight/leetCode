package same_tree

import (
	"github.com/leetCode/2021/common"
	"testing"
)

// [4 3 2 1 2 3 4]
func Test_isBalanced(t *testing.T) {
	n15 := &common.TreeNode{Val: 15}
	n7 := &common.TreeNode{Val: 7}
	n20 := &common.TreeNode{Val: 20, Left: n15, Right: n7}
	n9 := &common.TreeNode{Val: 9}
	n3 := &common.TreeNode{Val: 3, Left: n9, Right: n20}
	t.Log(isBalanced(n3)) // true
}

func Test_isBalanced2(t *testing.T) {
	n41 := &common.TreeNode{Val: 4}
	n42 := &common.TreeNode{Val: 4}
	n31 := &common.TreeNode{Val: 3, Left: n41, Right: n42}
	n32 := &common.TreeNode{Val: 3}
	n21 := &common.TreeNode{Val: 2, Left: n31, Right: n32}
	n22 := &common.TreeNode{Val: 2}
	n1 := &common.TreeNode{Val: 1, Left: n21, Right: n22}
	t.Log(isBalanced(n1)) // false
}

func Test_isBalanced3(t *testing.T) {
	n2, n3 := &common.TreeNode{Val: 2}, &common.TreeNode{Val: 3}
	n1 := &common.TreeNode{Val: 1, Right: n2}
	n2.Right = n3
	t.Log(isBalanced(n1)) // false
}

func Test_isBalanced4(t *testing.T) {
	n1, n21, n22 := &common.TreeNode{Val: 1}, &common.TreeNode{Val: 2}, &common.TreeNode{Val: 2}
	n31, n32, n41, n42 := &common.TreeNode{Val: 3}, &common.TreeNode{Val: 3}, &common.TreeNode{Val: 4}, &common.TreeNode{Val: 4}
	n1.Left, n1.Right, n21.Left, n22.Right = n21, n22, n31, n32
	n31.Left, n32.Right = n41, n42
	t.Log(isBalanced(n1)) // false
}
