package remove_duplicates_from_sorted_list_ii

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var top, p1 *ListNode
	p2 := head
	pre := head.Val - 1
	for p2 != nil {
		if p2.Val != pre {
			if (p2.Next != nil && p2.Val != p2.Next.Val) || p2.Next == nil {
				if top == nil {
					top = p2
					p1 = p2
				} else {
					p1.Next = p2
					p1 = p1.Next
				}
			}
			pre = p2.Val
		}
		p2 = p2.Next
		if p1 != nil {
			p1.Next = nil
		}
	}
	return top
}
