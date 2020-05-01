package binary_tree_right_side_view

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func RightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if i == l-1 {
				res = append(res, node.Val)
			}
		}
		queue = queue[l:]
	}
	return res
}
