package add_two_numbers

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	list1 := ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	list2 := ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	res := AddTwoNumbers(&list1, &list2)
	printListNode(res)

	list3 := ListNode{Val: 5}
	list4 := ListNode{Val: 5}
	res2 := AddTwoNumbers(&list3, &list4)
	printListNode(res2)
}
