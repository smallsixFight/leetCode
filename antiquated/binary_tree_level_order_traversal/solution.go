package binary_tree_level_order_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LevelOrderTwo(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		l := len(stack)
		arr := make([]int, 0)
		for i := 0; i < l; i++ {
			arr = append(arr, stack[i].Val)
			if stack[i].Left != nil {
				stack = append(stack, stack[i].Left)
			}
			if stack[i].Right != nil {
				stack = append(stack, stack[i].Right)
			}
		}
		res = append(res, arr)
		stack = stack[l:]
	}
	return res
}

func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	return dfsVisit(root, 0, make([][]int, 0))
}

func dfsVisit(node *TreeNode, depth int, record [][]int) [][]int {
	if node != nil {
		if len(record) < depth+1 {
			arr := make([]int, 0)
			record = append(record, arr)
		}
		record[depth] = append(record[depth], node.Val)
		record = dfsVisit(node.Left, depth+1, record)
		record = dfsVisit(node.Right, depth+1, record)
	}
	return record
}
