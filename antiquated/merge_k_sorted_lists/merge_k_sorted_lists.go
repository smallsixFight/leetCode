package merge_k_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeKLists(lists []*ListNode) *ListNode {
	if len(lists) < 1 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	l1, l2 := lists[0], lists[1]
	if len(lists) > 2 {
		mid := len(lists)/2 + 1
		l1 = MergeKLists(lists[0:mid])
		l2 = MergeKLists(lists[mid:])
	}
	top := &ListNode{}
	p := top
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				p.Next = l1
				l1 = l1.Next
			} else {
				p.Next = l2
				l2 = l2.Next
			}
		} else if l1 != nil {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	return top.Next
}

func MergeKLists2(lists []*ListNode) *ListNode {
	if len(lists) < 1 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	l1, l2 := lists[0], lists[1]
	if len(lists) > 2 {
		l1 = MergeKLists(lists[0:2])
		l2 = MergeKLists(lists[2:])
	}
	top := &ListNode{}
	p := top
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				p.Next = l1
				l1 = l1.Next
			} else {
				p.Next = l2
				l2 = l2.Next
			}
		} else if l1 != nil {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	return top.Next
}
