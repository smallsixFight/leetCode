package remove_duplicates_from_sorted_list

type ListNode struct {
	Val  int
	Next *ListNode
}

// My Version
func DeleteDuplicates2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
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

func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p1 := head
	for p1.Next != nil {
		if p1.Val == p1.Next.Val {
			p1.Next = p1.Next.Next
		} else {
			p1 = p1.Next
		}
	}
	return head
}
