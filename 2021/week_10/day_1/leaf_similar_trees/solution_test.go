package leaf_similar_trees

import (
	"github.com/leetCode/2021/common"
	"testing"
)

func Test_leafSimilar(t *testing.T) {
	root1 := &common.TreeNode{Val: 1}
	root2 := &common.TreeNode{Val: 2}
	t.Log(leafSimilar(root1, root2)) // false
}

func Test_leafSimilar2(t *testing.T) {
	t.Log(2 ^ 4 ^ 1)
	root1 := &common.TreeNode{Val: 1}
	root2 := &common.TreeNode{Val: 1}
	t.Log(leafSimilar(root1, root2)) // true
}
