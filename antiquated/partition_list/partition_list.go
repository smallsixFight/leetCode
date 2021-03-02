package partition_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	top := &ListNode{Next: head}
	pre, p := top, top
	for p != nil && p.Next != nil {
		if p.Next.Val >= x && pre == p {
			pre = p
		} else if p.Next.Val < x {
			if pre != p {
				temp := pre.Next
				pre.Next = p.Next
				p.Next = p.Next.Next
				pre.Next.Next = temp
				pre = pre.Next
				continue
			}
			pre = pre.Next
		}
		p = p.Next
	}
	return top.Next
}
