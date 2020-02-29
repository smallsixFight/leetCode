package remove_nth_node_from_end_of_list

import "testing"

func TestRemoveTthNodeFromEndOfListTwo2(t *testing.T) {
	head := ListNode{}
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5}
	head.Next = n1
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	res := RemoveTthNodeFromEndOfListTwo(n1, 2)
	for res != nil {
		t.Log(res.Val)
		res = res.Next
	}
}

func TestRemoveTthNodeFromEndOfListTwo(t *testing.T) {
	head := ListNode{Val: 1, Next: &ListNode{Val: 2}}
	res := RemoveTthNodeFromEndOfListTwo(&head, 2)
	for res != nil {
		t.Log(res.Val)
		res = res.Next
	}
}

func TestRemoveTthNodeFromEndOfList(t *testing.T) {
	head := ListNode{}
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5}
	head.Next = n1
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	res := RemoveTthNodeFromEndOfList(n1, 2)
	for res != nil {
		t.Log(res.Val)
		res = res.Next
	}
}

func TestRemoveTthNodeFromEndOfList2(t *testing.T) {
	head := ListNode{Val: 1}
	res := RemoveTthNodeFromEndOfList(&head, 1)
	for res != nil {
		t.Log(res.Val)
		res = res.Next
	}
}
