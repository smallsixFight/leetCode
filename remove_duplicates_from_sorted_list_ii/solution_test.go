package remove_duplicates_from_sorted_list_ii

import "testing"

func TestDeleteDuplicates(t *testing.T) {
	n5 := &ListNode{Val: 5}
	n44 := &ListNode{Val: 4, Next: n5}
	n4 := &ListNode{Val: 4, Next: n44}
	n33 := &ListNode{Val: 3, Next: n4}
	n3 := &ListNode{Val: 3, Next: n33}
	n2 := &ListNode{Val: 2, Next: n3}
	n1 := &ListNode{Val: 1, Next: n2}
	res := deleteDuplicates(n1)
	for res != nil {
		t.Log(res.Val, " ")
		res = res.Next
	}
}

func TestDeleteDuplicates2(t *testing.T) {
	n3 := &ListNode{Val: 3}
	n2 := &ListNode{Val: 2, Next: n3}
	n111 := &ListNode{Val: 1, Next: n2}
	n11 := &ListNode{Val: 1, Next: n111}
	n1 := &ListNode{Val: 1, Next: n11}
	res := deleteDuplicates(n1)
	for res != nil {
		t.Log(res.Val, " ")
		res = res.Next
	}
}
