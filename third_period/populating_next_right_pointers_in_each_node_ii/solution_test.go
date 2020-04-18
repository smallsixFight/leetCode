package populating_next_right_pointers_in_each_node_ii

import "testing"

func TestConnect(t *testing.T) {
	n4 := &Node{Val: 4}
	n5 := &Node{Val: 5}
	n7 := &Node{Val: 7}
	n2 := &Node{Val: 2, Left: n4, Right: n5}
	n3 := &Node{Val: 3, Right: n7}
	n1 := &Node{Val: 1, Left: n2, Right: n3}
	root := Connect(n1)
	root = Connect(root)
	for root != nil {
		p := root
		for p != nil {
			t.Log(p.Val)
			p = p.Next
		}
		root = root.Left
	}
}
