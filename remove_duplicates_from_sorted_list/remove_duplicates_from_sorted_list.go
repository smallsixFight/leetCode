package remove_duplicates_from_sorted_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteDuplicates(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	p1 := head.Next
	for p1.Next != nil {
		if p1.Val == p1.Next.Val {
			p1.Next = p1.Next.Next
		} else {
			p1 = p1.Next
		}
	}
	return head
}
