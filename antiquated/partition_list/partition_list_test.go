package partition_list

import "testing"

func TestPartition(t *testing.T) {
	n2 := &ListNode{Val: 2}
	n5 := &ListNode{Val: 5, Next: n2}
	n22 := &ListNode{Val: 2, Next: n5}
	n3 := &ListNode{Val: 3, Next: n22}
	n4 := &ListNode{Val: 4, Next: n3}
	n1 := &ListNode{Val: 1, Next: n4}
	res := partition(n1, 3)
	for res != nil {
		t.Log(res.Val, " ")
		res = res.Next
	}
}

func TestPartition2(t *testing.T) {
	n2 := &ListNode{Val: 2}
	n1 := &ListNode{Val: 1, Next: n2}
	n3 := &ListNode{Val: 3, Next: n1}
	res := partition(n3, 3)
	for res != nil {
		t.Log(res.Val, " ")
		res = res.Next
	}
}
