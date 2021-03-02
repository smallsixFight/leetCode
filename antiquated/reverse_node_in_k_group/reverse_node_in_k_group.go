package reverse_node_in_k_group

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseKGroup(head *ListNode, k int) *ListNode {
	top := &ListNode{Next: head}
	pre, end := top, top
	for end != nil {
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		if end == nil {
			break
		}
		start := pre.Next
		next := end.Next
		end.Next = nil
		pre.Next = reverseList(start)
		start.Next = next
		pre = start
		end = pre
	}
	return top.Next
}

func reverseList(list *ListNode) *ListNode {
	var pre, next *ListNode
	curr := list
	for curr != nil {
		next = curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}
	return pre
}
