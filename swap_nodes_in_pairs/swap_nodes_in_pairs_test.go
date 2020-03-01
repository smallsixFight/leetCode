package swap_nodes_in_pairs

import (
	"fmt"
	"testing"
)

func TestSwapNodesInPairs(t *testing.T) {
	n1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}
	res := SwapNodesInPairs(n1)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
