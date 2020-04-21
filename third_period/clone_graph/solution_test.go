package clone_graph

import "testing"

func TestCloneGraph(t *testing.T) {
	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n3 := &Node{Val: 3}
	n4 := &Node{Val: 4}
	n1.Neighbors = []*Node{n2, n4}
	n3.Neighbors = []*Node{n2, n4}
	n2.Neighbors = []*Node{n1, n3}
	n4.Neighbors = []*Node{n1, n3}
	t.Log(n1, n2, n3, n4)
	res := CloneGraph(n1)
	t.Log(res, res.Neighbors[0], res.Neighbors[1].Neighbors[1], res.Neighbors[1].Neighbors[1].Neighbors[1])
}
