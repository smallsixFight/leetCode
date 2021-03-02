package reverse_linked_list_ii

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseBetween(head *ListNode, m, n int) *ListNode {
	if head == nil || head.Next == nil || m >= n {
		return head
	}
	top := &ListNode{Next: head}
	pre, sp := top, top
	s, p := head, head
	count := 1
	for p != nil {
		if count < m {
			s = s.Next
			sp = sp.Next
			pre = pre.Next
		} else if count == m {
			pre = pre.Next
		} else if count <= n {
			temp := p.Next
			p.Next = pre
			pre = p
			if count == n {
				sp.Next = p
				s.Next = temp
				break
			}
			p = temp
			count++
			continue
		}
		count++
		p = p.Next
	}
	return top.Next
}
