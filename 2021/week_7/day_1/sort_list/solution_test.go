package sort_list

import "testing"

func Test_sortList2(t *testing.T) {
	n3 := &ListNode{Val: 3}
	n1 := &ListNode{Val: 1, Next: n3}
	n2 := &ListNode{Val: 2, Next: n1}
	n4 := &ListNode{Val: 4, Next: n2}
	res := sortList2(n4)
	for res != nil {
		t.Log(res.Val)
		res = res.Next
	}
}

func Test_sortList2_2(t *testing.T) {
	n0 := &ListNode{Val: 0}
	n4 := &ListNode{Val: 4, Next: n0}
	n3 := &ListNode{Val: 3, Next: n4}
	n5 := &ListNode{Val: 5, Next: n3}
	n1 := &ListNode{Val: -1, Next: n5}
	res := sortList2(n1)
	for res != nil {
		t.Log(res.Val)
		res = res.Next
	}
}

func Test_sortList(t *testing.T) {
	n3 := &ListNode{Val: 3}
	n1 := &ListNode{Val: 1, Next: n3}
	n2 := &ListNode{Val: 2, Next: n1}
	n4 := &ListNode{Val: 4, Next: n2}
	res := sortList(n4)
	for res != nil {
		t.Log(res.Val)
		res = res.Next
	}
}

func Test_sortList_2(t *testing.T) {
	n0 := &ListNode{Val: 0}
	n4 := &ListNode{Val: 4, Next: n0}
	n3 := &ListNode{Val: 3, Next: n4}
	n5 := &ListNode{Val: 5, Next: n3}
	n1 := &ListNode{Val: -1, Next: n5}
	res := sortList(n1)
	for res != nil {
		t.Log(res.Val)
		res = res.Next
	}
}
