package populating_next_right_pointers_in_each_node_ii

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
	if node.Left == nil && node.Right == nil {
		return
	}
	if node.Right != nil {
		if uncleNode != nil {
			if uncleNode.Left != nil {
				node.Right.Next = uncleNode.Left
			} else {
				node.Right.Next = uncleNode.Right
			}
		}
		temp := node.Right.Next
		for temp != nil && temp.Left == nil && temp.Right == nil {
			temp = temp.Next
		}
		dfs(node.Right, temp)
	}
	if node.Left != nil {
		if node.Right != nil {
			node.Left.Next = node.Right
		} else if uncleNode != nil {
			if uncleNode.Left != nil {
				node.Left.Next = uncleNode.Left
			} else {
				node.Left.Next = uncleNode.Right
			}
		}
		temp := node.Left.Next
		for temp != nil && temp.Left == nil && temp.Right == nil {
			temp = temp.Next
		}
		dfs(node.Left, temp)
	}
}
