package binary_tree_level_order_traversal_II

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LevelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	record := make([][]int, 0)
	record = dfsVisit(root, 0, record)
	// 逆序
	for i, j := 0, len(record)-1; i < j; i++ {
		record[i], record[j] = record[j], record[i]
		j--
	}
	return record
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
