package swap_nodes_in_pairs

type ListNode struct {
	Val  int
	Next *ListNode
}

func SwapNodesInPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	top := &ListNode{Next: head}
	pre := top
	for pre.Next != nil {
		if pre.Next.Next != nil {
			p1, p2 := pre.Next, pre.Next.Next
			pre.Next = p2
			p1.Next = p2.Next
			p2.Next = p1
			pre = pre.Next.Next
		} else {
			break
		}
	}
	return top.Next
}
