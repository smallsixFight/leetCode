package clone_graph

type Node struct {
	Val       int
	Neighbors []*Node
}

func CloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	hashMap := make(map[*Node]*Node)
	root := &Node{Val: node.Val, Neighbors: make([]*Node, 0)}
	hashMap[node] = root
	queue := []*Node{node}
	for len(queue) > 0 {
		n := queue[0]
		for i := 0; i < len(n.Neighbors); i++ {
			if _, ok := hashMap[n.Neighbors[i]]; !ok {
				newNode := &Node{Val: n.Neighbors[i].Val, Neighbors: make([]*Node, 0)}
				hashMap[n].Neighbors = append(hashMap[n].Neighbors, newNode)
				hashMap[n.Neighbors[i]] = newNode
				queue = append(queue, n.Neighbors[i])
			} else {
				hashMap[n].Neighbors = append(hashMap[n].Neighbors, hashMap[n.Neighbors[i]])
			}
		}
		queue = queue[1:]
	}
	return root
}
