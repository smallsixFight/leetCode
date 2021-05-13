package leaf_similar_trees

import (
	"github.com/leetCode/2021/common"
	"testing"
)

func Test_flatten(t *testing.T) {
	n3 := &common.TreeNode{Val: 3}
	n4 := &common.TreeNode{Val: 4}
	n6 := &common.TreeNode{Val: 6}
	n2 := &common.TreeNode{Val: 2, Left: n3, Right: n4}
	n5 := &common.TreeNode{Val: 5, Right: n6}
	n1 := &common.TreeNode{Val: 1, Left: n2, Right: n5}
	flatten(n1)
	for n1 != nil {
		t.Log(n1.Val)
		n1 = n1.Right
	}
}
