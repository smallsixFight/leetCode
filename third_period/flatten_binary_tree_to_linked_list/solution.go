package flatten_binary_tree_to_linked_list

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Flatten(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root)
	root.Left = nil
}

func dfs(node *TreeNode) *TreeNode {
	var p1, rp *TreeNode
	if node.Left != nil && node.Right != nil {
		rp = node.Right
	}
	if node.Left != nil {
		node.Right = node.Left
		p1 = dfs(node.Left)
	} else if node.Right != nil { // 如果左孩子为空，直接对右子树进行 DFS
		dfs(node.Right)
	}
	if rp != nil { // 左子树不为空且右子树不为空的情况
		if p1 != nil {
			for p1.Right != nil {
				p1 = p1.Right
			}
			p1.Right = rp
			dfs(rp)
		} else {
			node.Left.Right = rp
		}
	}
	node.Left = nil // 清空左孩子的引用
	return node
}
