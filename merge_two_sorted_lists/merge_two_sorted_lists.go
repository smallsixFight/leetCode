package merge_two_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoSortedLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head, pre := &ListNode{}, &ListNode{}
	pre = head
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	for l1 != nil || l2 != nil {
		if l1 == nil {
			pre.Next = l2
			l2 = l2.Next
		} else if l2 == nil {
			pre.Next = l1
			l1 = l1.Next
		} else {
			if l1.Val < l2.Val {
				pre.Next = l1
				l1 = l1.Next
			} else {
				pre.Next = l2
				l2 = l2.Next
			}
		}
		pre = pre.Next
	}
	return head.Next
}
