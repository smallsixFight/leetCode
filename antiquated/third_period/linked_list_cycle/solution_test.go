package linked_list_cycle

import (
	"testing"
)

func TestHasCycle(t *testing.T) {
	n4 := &ListNode{Val: 4}
	n0 := &ListNode{Val: 0, Next: n4}
	n2 := &ListNode{Val: 2, Next: n0}
	n3 := &ListNode{Val: 3, Next: n2}
	n4.Next = n2
	t.Log(HasCycle(n3))
}
