package merge_two_sorted_lists

import "testing"

func TestMergeTwoSortedLists(t *testing.T) {
	l1 := ListNode{Val: 1}
	l1.Next = &ListNode{Val: 2}
	l1.Next.Next = &ListNode{Val: 4}

	l2 := ListNode{Val: 1}
	l2.Next = &ListNode{Val: 3}
	l2.Next.Next = &ListNode{Val: 4}

	result := MergeTwoSortedLists(&l1, &l2)
	for result != nil {
		t.Log(result.Val)
		result = result.Next
	}
}
