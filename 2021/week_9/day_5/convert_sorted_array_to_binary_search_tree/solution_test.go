package same_tree

import (
	"github.com/leetCode/2021/common"
	"testing"
)

func Test_sortedArrayToBST(t *testing.T) {
	root := sortedArrayToBST([]int{-10, -3, 0, 5, 9})
	t.Log(common.GetTreeStrings(root))
	root = sortedArrayToBST([]int{1, 3})
	t.Log(common.GetTreeStrings(root))
}
