package construct_binary_tree_from_preorder_and_inorder_traversal

import (
	"github.com/leetCode/2021/common"
	"testing"
)

func Test_buildTree(t *testing.T) {
	tree := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	t.Log(common.GetTreeStrings(tree))
}
