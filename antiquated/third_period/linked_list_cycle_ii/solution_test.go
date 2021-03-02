package linked_list_cycle_ii

import "testing"

func TestDetectCycleThree(t *testing.T) {
	n4 := &ListNode{Val: 4}
	n0 := &ListNode{Val: 0, Next: n4}
	n2 := &ListNode{Val: 2, Next: n0}
	n3 := &ListNode{Val: 3, Next: n2}
	n4.Next = n2
	t.Log(DetectCycleThree(n3).Val)
}

func TestDetectCycleTwo(t *testing.T) {
	n4 := &ListNode{Val: 4}
	n0 := &ListNode{Val: 0, Next: n4}
	n2 := &ListNode{Val: 2, Next: n0}
	n3 := &ListNode{Val: 3, Next: n2}
	n4.Next = n2
	t.Log(DetectCycleTwo(n3).Val)
}

func TestDetectCycle(t *testing.T) {
	n4 := &ListNode{Val: 4}
	n0 := &ListNode{Val: 0, Next: n4}
	n2 := &ListNode{Val: 2, Next: n0}
	n3 := &ListNode{Val: 3, Next: n2}
	n4.Next = n2
	t.Log(DetectCycle(n3).Val)
}
