package merge_k_sorted_lists

import (
	"fmt"
	"testing"
)

func TestMergeKLists_2(t *testing.T) {
	lists := make([]*ListNode, 3)
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	l3 := &ListNode{Val: 2, Next: &ListNode{Val: 6}}
	lists[0] = l1
	lists[1] = l2
	lists[2] = l3
	res := MergeKLists2(lists)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

func TestMergeKLists(t *testing.T) {
	lists := make([]*ListNode, 3)
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	l3 := &ListNode{Val: 2, Next: &ListNode{Val: 6}}
	lists[0] = l1
	lists[1] = l2
	lists[2] = l3
	res := MergeKLists(lists)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
