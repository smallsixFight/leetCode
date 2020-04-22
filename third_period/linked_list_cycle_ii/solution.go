package linked_list_cycle_ii

type ListNode struct {
	Val  int
	Next *ListNode
}

func DetectCycleThree(head *ListNode) *ListNode {
	p1, p2 := head, head
	for p2 != nil && p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
		if p1 == p2 {
			p1 = head
			for p1 != p2 {
				p1 = p1.Next
				p2 = p2.Next
			}
			return p1
		}
	}
	return nil
}

func DetectCycleTwo(head *ListNode) *ListNode {
	m := make(map[*ListNode]bool)
	for head != nil {
		if m[head] {
			return head
		}
		m[head] = true
		head = head.Next
	}
	return nil
}

func DetectCycle(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	p1, p2 := head, head.Next
	for p1 != p2 {
		if p2 != nil && p2.Next != nil {
			p2 = p2.Next.Next
			p1 = p1.Next
		} else {
			return nil
		}
	}
	p1 = head
	p3 := p2
	for p1 != p2 {
		if p3 == p1 {
			return p1
		}
		p3 = p3.Next
		if p3 == p2 {
			p1 = p1.Next
		}
	}
	return p1
}
