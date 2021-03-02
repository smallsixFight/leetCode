package intersection_of_two_linked_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func GetIntersectionNodeTwo(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	p1, p2 := headA, headB
	move1, move2 := false, false
	for p1 != nil && p2 != nil {
		if p1 == p2 {
			return p1
		}
		if p1.Next == nil && !move1 {
			p1 = headB
			move1 = true
		} else {
			p1 = p1.Next
		}
		if p2.Next == nil && !move2 {
			move2 = true
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}
	return nil
}

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	p1, p2 := headA, headB
	l1, l2 := 0, 0
	for p1 != nil || p2 != nil {
		if p1 != nil {
			l1++
			p1 = p1.Next
		}
		if p2 != nil {
			l2++
			p2 = p2.Next
		}
	}
	p1, p2 = headA, headB
	for l1 > l2 {
		l1--
		p1 = p1.Next
	}
	for l2 > l1 {
		l2--
		p2 = p2.Next
	}
	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p1
}
