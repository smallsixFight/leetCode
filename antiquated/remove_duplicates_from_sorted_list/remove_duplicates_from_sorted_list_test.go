package remove_duplicates_from_sorted_list

import "testing"

func TestDeleteDuplicates2(t *testing.T) {
	head := &ListNode{}
	n0 := ListNode{Val: 1}
	n1 := ListNode{Val: 1}
	n2 := ListNode{Val: 2}
	head.Next = &n0
	n0.Next = &n1
	n1.Next = &n2
	res := DeleteDuplicates2(head).Next
	for res != nil {
		t.Log(res.Val)
		res = res.Next
	}
	t.Log("------------------------------------------")
	n3 := ListNode{Val: 1}
	n4 := ListNode{Val: 1}
	n5 := ListNode{Val: 2}
	n6 := ListNode{Val: 3}
	n7 := ListNode{Val: 3}
	head1 := &ListNode{}
	head1.Next = &n3
	n3.Next = &n4
	n4.Next = &n5
	n5.Next = &n6
	n6.Next = &n7
	res = DeleteDuplicates2(head1).Next
	for res != nil {
		t.Log(res.Val)
		res = res.Next
	}
}

func TestDeleteDuplicates(t *testing.T) {
	head := ListNode{}
	n0 := ListNode{Val: 1}
	n1 := ListNode{Val: 1}
	n2 := ListNode{Val: 2}
	head.Next = &n0
	n0.Next = &n1
	n1.Next = &n2
	res := DeleteDuplicates(&head).Next
	for res != nil {
		t.Log(res.Val)
		res = res.Next
	}
	t.Log("------------------------------------------")
	n3 := ListNode{Val: 1}
	n4 := ListNode{Val: 1}
	n5 := ListNode{Val: 2}
	n6 := ListNode{Val: 3}
	n7 := ListNode{Val: 3}
	head1 := ListNode{}
	head1.Next = &n3
	n3.Next = &n4
	n4.Next = &n5
	n5.Next = &n6
	n6.Next = &n7
	res = DeleteDuplicates(&head1).Next
	for res != nil {
		t.Log(res.Val)
		res = res.Next
	}
}
