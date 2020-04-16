package path_sum_ii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res [][]int

func PathSum(root *TreeNode, sum int) [][]int {
	if root == nil {

	}
	res = make([][]int, 0)
	dfs(root, sum, make([]int, 0))
	return res
}

func dfs(node *TreeNode, sum int, arr []int) {
	if node == nil || (node.Val != sum && node.Left == nil && node.Right == nil) {
		return
	}
	arr = append(arr, node.Val)
	if node.Val == sum && node.Left == nil && node.Right == nil {
		res = append(res, arr)
		return
	}
	if node.Left != nil || node.Right != nil {
		if node.Left != nil && node.Right != nil {
			dfs(node.Left, sum-node.Val, arr)
			newArr := make([]int, len(arr))
			copy(newArr, arr)
			dfs(node.Right, sum-node.Val, newArr)
		} else if node.Left != nil {
			dfs(node.Left, sum-node.Val, arr)
		} else {
			dfs(node.Right, sum-node.Val, arr)
		}
	}
}
