package reorder_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReorderList(head *ListNode) {
	if head == nil {
		return
	}
	arr := make([]*ListNode, 0)
	for head != nil {
		arr = append(arr, head)
		head = head.Next
	}
	p1, p2 := 0, len(arr)-1
	for p1 < p2 {
		arr[p1].Next = arr[p2]
		arr[p2].Next = arr[p1+1]
		p1++
		p2--
	}
	arr[p1].Next = nil
}
