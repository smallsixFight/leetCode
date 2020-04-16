package convert_sorted_list_to_binary_search_tree

import "testing"

func TestSortedListToBFT(t *testing.T) {
	// -10,-3,0,5,9
	list := &ListNode{Val: -10, Next: &ListNode{Val: -3,
		Next: &ListNode{Val: 0, Next: &ListNode{Val: 5, Next: &ListNode{Val: 9}}}}}
	t.Log(getTreeArr(SortedListToBFT(list)))
}
