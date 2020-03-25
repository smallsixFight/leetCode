package rotate_list

import (
	"fmt"
	"testing"
)

func TestRotateRightTwo(t *testing.T) {
	n5 := &ListNode{Val: 5}
	n4 := &ListNode{Val: 4, Next: n5}
	n3 := &ListNode{Val: 3, Next: n4}
	n2 := &ListNode{Val: 2, Next: n3}
	n1 := &ListNode{Val: 1, Next: n2}
	res := RotateRightTwo(n1, 2)
	for res != nil {
		fmt.Print(res.Val, " ")
		res = res.Next
	}
	fmt.Println()
	n0 := &ListNode{Val: 0, Next: n1}
	n1.Val = 1
	n2.Val = 2
	n2.Next = nil
	res = RotateRightTwo(n0, 4)
	for res != nil {
		fmt.Print(res.Val, " ")
		res = res.Next
	}
	fmt.Println()
}

func TestRotateRight(t *testing.T) {
	n5 := &ListNode{Val: 5}
	n4 := &ListNode{Val: 4, Next: n5}
	n3 := &ListNode{Val: 3, Next: n4}
	n2 := &ListNode{Val: 2, Next: n3}
	n1 := &ListNode{Val: 1, Next: n2}
	res := RotateRight(n1, 2)
	for res != nil {
		fmt.Print(res.Val, " ")
		res = res.Next
	}
	fmt.Println()
	n0 := &ListNode{Val: 0, Next: n1}
	n1.Val = 1
	n2.Val = 2
	n2.Next = nil
	res = RotateRight(n0, 4)
	for res != nil {
		fmt.Print(res.Val, " ")
		res = res.Next
	}
	fmt.Println()

}
