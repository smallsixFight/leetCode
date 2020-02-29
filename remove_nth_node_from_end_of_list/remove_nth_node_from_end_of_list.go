package remove_nth_node_from_end_of_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveTthNodeFromEndOfListTwo(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	var pre *ListNode
	r, count := head, 1
	for r.Next != nil {
		r = r.Next
		if count < n {
			count++
		} else if pre == nil {
			pre = head
		} else {
			pre = pre.Next
		}
	}
	if pre != nil {
		pre.Next = pre.Next.Next
	} else {
		head = head.Next
	}
	return head
}
func RemoveTthNodeFromEndOfList(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	arr := make([]*ListNode, 0)
	p := head
	for p != nil {
		arr = append(arr, p)
		p = p.Next
	}
	idx := len(arr) - 1 - n
	if idx < 0 {
		head = head.Next
	} else {
		arr[idx].Next = arr[idx].Next.Next
	}
	return head
}
