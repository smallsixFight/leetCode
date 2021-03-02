package rotate_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func RotateRightTwo(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	l, p := 0, head
	for p != nil {
		l++
		p = p.Next
	}
	k %= l
	if k == 0 {
		return head
	}
	target := l + 1 - k
	top := ListNode{Next: head}
	p = head

	for i := 1; ; i++ {
		if i == target-1 {
			p1 := p
			p = p.Next
			p1.Next = nil
			continue
		} else if i == target {
			top.Next = p
		}
		if p.Next == nil {
			p.Next = head
			break
		}
		p = p.Next
	}
	return top.Next
}

func RotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	l, p := 0, head
	record, NewRecord := make(map[int]int), make(map[int]int)
	for p != nil {
		l++
		record[l] = p.Val
		p = p.Next
	}
	k %= l
	if k == 0 {
		return head
	}
	for i := 1; i <= l; i++ {
		idx := i + k
		if idx > l {
			idx -= l
		}
		NewRecord[idx] = record[i]
	}
	p = head
	for idx := 1; p != nil; idx++ {
		p.Val = NewRecord[idx]
		p = p.Next
	}
	return head
}
