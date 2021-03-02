package insertion_sort_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func InsertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	top := &ListNode{Next: head}
	p1, p2 := head, head.Next
	for p2 != nil {
		if p2.Val < p1.Val {
			temp := p2.Next
			p1.Next = p2.Next
			p3 := top
			for p3.Next.Val < p2.Val {
				p3 = p3.Next
			}
			p2.Next = p3.Next
			p3.Next = p2
			p2 = temp
		} else {
			p1 = p1.Next
			p2 = p2.Next
		}
	}
	return top.Next
}
