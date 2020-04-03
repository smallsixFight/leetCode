package reverse_linked_list_ii

import (
	"fmt"
	"testing"
)

func TestReverseBetween2(t *testing.T) {
	n3 := &ListNode{Val: 3}
	n2 := &ListNode{Val: 2, Next: n3}
	n1 := &ListNode{Val: 1, Next: n2}
	res := ReverseBetween(n1, 2, 3)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

func TestReverseBetween(t *testing.T) {
	n5 := &ListNode{Val: 5}
	n4 := &ListNode{Val: 4, Next: n5}
	n3 := &ListNode{Val: 3, Next: n4}
	n2 := &ListNode{Val: 2, Next: n3}
	n1 := &ListNode{Val: 1, Next: n2}
	res := ReverseBetween(n1, 2, 4)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
