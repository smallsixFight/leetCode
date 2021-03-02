package binary_tree_zigzag_level_order_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ZigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	record := []*TreeNode{root}
	isReverse := false
	for len(record) > 0 {
		arr := make([]int, 0)
		l := len(record)
		//if isReverse {
		for i := l - 1; i >= 0; i-- {
			arr = append(arr, record[i].Val)

			if isReverse {
				if record[i].Right != nil {
					record = append(record, record[i].Right)
				}
				if record[i].Left != nil {
					record = append(record, record[i].Left)
				}
			} else {
				if record[i].Left != nil {
					record = append(record, record[i].Left)
				}
				if record[i].Right != nil {
					record = append(record, record[i].Right)
				}
			}
		}
		res = append(res, arr)
		record = record[l:]
		isReverse = !isReverse
	}
	return res
}
