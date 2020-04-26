package sort_list

import "testing"

func TestSortListWithMergeSort2(t *testing.T) {
	n3 := &ListNode{Val: 3}
	n1 := &ListNode{Val: 1, Next: n3}
	n2 := &ListNode{Val: 2, Next: n1}
	n4 := &ListNode{Val: 4, Next: n2}
	res := SortListWithMergeSort(n4)
	for res != nil {
		t.Log(res.Val)
		res = res.Next
	}
}

func TestSortListWithMergeSort(t *testing.T) {
	n0 := &ListNode{Val: 0}
	n4 := &ListNode{Val: 4, Next: n0}
	n3 := &ListNode{Val: 3, Next: n4}
	n5 := &ListNode{Val: 5, Next: n3}
	n1 := &ListNode{Val: -1, Next: n5}
	res := SortListWithMergeSort(n1)
	for res != nil {
		t.Log(res.Val)
		res = res.Next
	}
}

func TestSortListWithQuickSort2(t *testing.T) {
	n0 := &ListNode{Val: 0}
	n4 := &ListNode{Val: 4, Next: n0}
	n3 := &ListNode{Val: 3, Next: n4}
	n5 := &ListNode{Val: 5, Next: n3}
	n1 := &ListNode{Val: -1, Next: n5}
	res := SortListWithQuickSort(n1)
	for res != nil {
		t.Log(res.Val)
		res = res.Next
	}
}

func TestSortListWithQuickSort(t *testing.T) {
	n3 := &ListNode{Val: 3}
	n1 := &ListNode{Val: 1, Next: n3}
	n2 := &ListNode{Val: 2, Next: n1}
	n4 := &ListNode{Val: 4, Next: n2}
	res := SortListWithQuickSort(n4)
	for res != nil {
		t.Log(res.Val)
		res = res.Next
	}
}
