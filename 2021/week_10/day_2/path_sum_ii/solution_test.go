package leaf_similar_trees

import (
	"github.com/leetCode/2021/common"
	"testing"
)

func Test_pathSum(t *testing.T) {
	n7 := &common.TreeNode{Val: 7}
	n2 := &common.TreeNode{Val: 2}
	n51 := &common.TreeNode{Val: 5}
	n1 := &common.TreeNode{Val: 1}
	n11 := &common.TreeNode{Val: 11, Left: n7, Right: n2}
	n41 := &common.TreeNode{Val: 4, Left: n51, Right: n1}
	n42 := &common.TreeNode{Val: 4, Left: n11}
	n13 := &common.TreeNode{Val: 13}
	n8 := &common.TreeNode{Val: 8, Left: n13, Right: n41}
	n52 := &common.TreeNode{Val: 5, Left: n42, Right: n8}
	t.Log(pathSum(n52, 22))
}
