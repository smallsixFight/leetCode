package same_tree

import (
	"github.com/leetCode/2021/common"
	"testing"
)

func Test_minDepth(t *testing.T) {
	n15 := &common.TreeNode{Val: 15}
	n7 := &common.TreeNode{Val: 7}
	n20 := &common.TreeNode{Val: 20, Left: n15, Right: n7}
	n9 := &common.TreeNode{Val: 9}
	n3 := &common.TreeNode{Val: 3, Left: n9, Right: n20}
	t.Log(minDepth(n3)) // 2
}
