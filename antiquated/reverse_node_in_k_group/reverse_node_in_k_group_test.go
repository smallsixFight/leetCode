package reverse_node_in_k_group

import (
	"fmt"
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	list := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	res := ReverseKGroup(list, 2)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
