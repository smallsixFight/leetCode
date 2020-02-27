package add_two_numbers

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func printListNode(node *ListNode) {
	for node != nil {
		fmt.Print(node.Val, ", ")
		node = node.Next
	}
}

func AddTwoNumbers(l1, l2 *ListNode) *ListNode {
	head := &ListNode{}
	p := head
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		temp := 0 + carry
		if l1 != nil {
			temp += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			temp += l2.Val
			l2 = l2.Next
		}
		if temp > 9 {
			carry = 1
		} else {
			carry = 0
		}
		node := &ListNode{Val: temp % 10}
		p.Next = node
		p = p.Next
	}
	return head.Next
}
