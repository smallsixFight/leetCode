package reorder_list

import "testing"

func TestReorderList2(t *testing.T) {
	n5 := &ListNode{Val: 5}
	n4 := &ListNode{Val: 4, Next: n5}
	n3 := &ListNode{Val: 3, Next: n4}
	n2 := &ListNode{Val: 2, Next: n3}
	n1 := &ListNode{Val: 1, Next: n2}
	ReorderList(n1)
	for n1 != nil {
		t.Log(n1.Val)
		n1 = n1.Next
	}
}

func TestReorderList(t *testing.T) {
	n4 := &ListNode{Val: 4}
	n3 := &ListNode{Val: 3, Next: n4}
	n2 := &ListNode{Val: 2, Next: n3}
	n1 := &ListNode{Val: 1, Next: n2}
	ReorderList(n1)
	for n1 != nil {
		t.Log(n1.Val)
		n1 = n1.Next
	}
}
