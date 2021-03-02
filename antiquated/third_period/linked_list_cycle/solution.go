package linked_list_cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func HasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	p1, p2 := head, head.Next
	for p1 != p2 {
		if p2 != nil && p2.Next != nil {
			p2 = p2.Next.Next
		} else {
			return false
		}
		p1 = p1.Next
	}
	return true
}
