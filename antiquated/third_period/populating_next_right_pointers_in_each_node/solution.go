package populating_next_right_pointers_in_each_node

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func Connect(root *Node) *Node {
	if root == nil {
		return root
	}
	dfs(root, nil)
	return root
}

func dfs(node *Node, uncleNode *Node) {
	if node.Left == nil {
		return
	}
	node.Left.Next = node.Right
	if uncleNode != nil {
		node.Right.Next = uncleNode.Left
	}
	dfs(node.Left, node.Right)
	if node.Next != nil {
		dfs(node.Right, node.Next.Left)
	} else {
		dfs(node.Right, nil)
	}
}
