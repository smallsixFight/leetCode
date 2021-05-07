package same_tree

import (
	"github.com/leetCode/2021/common"
	"testing"
)

func Test_sortedListToBST2(t *testing.T) {
	list := &common.ListNode{Val: -10, Next: &common.ListNode{Val: -3,
		Next: &common.ListNode{Val: 0, Next: &common.ListNode{Val: 5, Next: &common.ListNode{Val: 9}}}}}
	root := sortedListToBST2(list)
	t.Log(common.GetTreeStrings(root))
}

func Test_sortedListToBST(t *testing.T) {
	list := &common.ListNode{Val: -10, Next: &common.ListNode{Val: -3,
		Next: &common.ListNode{Val: 0, Next: &common.ListNode{Val: 5, Next: &common.ListNode{Val: 9}}}}}
	root := sortedListToBST(list)
	t.Log(common.GetTreeStrings(root))
}
