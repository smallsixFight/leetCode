package populating_next_right_pointers_in_each_node

import "testing"

func TestConnect2(t *testing.T) {
	n6 := &Node{Val: 6}
	n7 := &Node{Val: 7}
	n8 := &Node{Val: 8}
	n9 := &Node{Val: 9}
	n10 := &Node{Val: 10}
	n11 := &Node{Val: 11}
	n12 := &Node{Val: 12}
	n13 := &Node{Val: 13}
	n2 := &Node{Val: 2, Left: n6, Right: n7}
	n3 := &Node{Val: 3, Left: n8, Right: n9}
	n4 := &Node{Val: 4, Left: n10, Right: n11}
	n5 := &Node{Val: 5, Left: n12, Right: n13}
	n0 := &Node{Val: 0, Left: n2, Right: n3}
	n1 := &Node{Val: 1, Left: n4, Right: n5}
	root := &Node{Val: -1, Left: n0, Right: n1}
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

func TestConnect(t *testing.T) {
	n4 := &Node{Val: 4}
	n5 := &Node{Val: 5}
	n6 := &Node{Val: 6}
	n7 := &Node{Val: 7}
	n2 := &Node{Val: 2, Left: n4, Right: n5}
	n3 := &Node{Val: 3, Left: n6, Right: n7}
	n1 := &Node{Val: 1, Left: n2, Right: n3}
	n1 = Connect(n1)
	for n1 != nil {
		p := n1
		for p != nil {
			t.Log(p.Val)
			p = p.Next
		}
		n1 = n1.Left
	}
}
