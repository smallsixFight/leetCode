package sort_colors

import "testing"

func Test_insertionSortList(t *testing.T) {
	n3 := &ListNode{Val: 3}
	n1 := &ListNode{Val: 1, Next: n3}
	n2 := &ListNode{Val: 2, Next: n1}
	n4 := &ListNode{Val: 4, Next: n2}
	res := insertionSortList(n4)
	for res != nil {
		t.Log(res.Val)
		res = res.Next
	}
}
