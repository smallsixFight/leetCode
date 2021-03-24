package next_greater_node_in_linked_list

import "testing"

func Test_nextLargerNodes(t *testing.T) {
	head := ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 5}}}
	t.Log(nextLargerNodes(&head))

	head = ListNode{Val: 2, Next: &ListNode{Val: 7, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}}}}
	t.Log(nextLargerNodes(&head))

	head = ListNode{Val: 1, Next: &ListNode{Val: 7, Next: &ListNode{Val: 5,
		Next: &ListNode{Val: 1, Next: &ListNode{Val: 9, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5,
			Next: &ListNode{Val: 1}}}}}}}}
	t.Log(nextLargerNodes(&head))
	t.Log(nextLargerNodes(nil))
}
