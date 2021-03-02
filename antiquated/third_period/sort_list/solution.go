package sort_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func SortListWithMergeSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre := &ListNode{}
	p1, p2 := head, head
	for p2 != nil && p2.Next != nil {
		pre = p1
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	l1 := SortListWithMergeSort(p1)
	pre.Next = nil
	l2 := SortListWithMergeSort(head)
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	top := &ListNode{Next: l1}
	p3 := top
	for l2 != nil {
		if p3.Next == nil {
			p3.Next = l2
			break
		} else if l2.Val < p3.Next.Val {
			temp := l2.Next
			l2.Next = p3.Next
			p3.Next = l2
			l2 = temp
		}
		p3 = p3.Next
	}
	return top.Next
}

func SortListWithQuickSort(head *ListNode) *ListNode {
	top := &ListNode{Next: head}
	sort(top, nil)
	return top.Next
}

func sort(top, end *ListNode) {
	if top.Next == nil || top.Next == end {
		return
	}
	n := top.Next
	e1 := top.Next
	for n.Next != end {
		if n.Next.Val < top.Next.Val {
			temp := n.Next
			n.Next = n.Next.Next
			temp.Next = top.Next
			top.Next = temp
		} else {
			n = n.Next
		}
	}
	sort(top, e1)
	sort(e1, nil)
}
