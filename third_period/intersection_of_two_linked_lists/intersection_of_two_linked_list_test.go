package intersection_of_two_linked_lists

import "testing"

func TestGetIntersectionNodeTwo(t *testing.T) {
	n6 := &ListNode{Val: 6}
	n5 := &ListNode{Val: 5, Next: n6}
	n4 := &ListNode{Val: 4, Next: n5}

	n11 := &ListNode{Val: 2, Next: n4}
	l1 := &ListNode{Val: 1, Next: n11}

	n22 := &ListNode{Val: 3, Next: n4}
	n21 := &ListNode{Val: 2, Next: n22}
	l2 := &ListNode{Val: 1, Next: n21}
	t.Log(GetIntersectionNodeTwo(l1, l2))
}

func TestGetIntersectionNode(t *testing.T) {
	n6 := &ListNode{Val: 6}
	n5 := &ListNode{Val: 5, Next: n6}
	n4 := &ListNode{Val: 4, Next: n5}

	n11 := &ListNode{Val: 2, Next: n4}
	l1 := &ListNode{Val: 1, Next: n11}

	n22 := &ListNode{Val: 3, Next: n4}
	n21 := &ListNode{Val: 2, Next: n22}
	l2 := &ListNode{Val: 1, Next: n21}
	t.Log(GetIntersectionNode(l1, l2))
}
