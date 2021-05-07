package construct_binary_tree_from_inorder_and_postorder_traversal

import (
	"github.com/leetCode/2021/common"
	"testing"
)

func Test_buildTree(t *testing.T) {
	tree := buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
	t.Log(common.GetTreeStrings(tree))
}
