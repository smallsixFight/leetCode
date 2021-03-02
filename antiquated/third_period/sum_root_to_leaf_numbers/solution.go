package sum_root_to_leaf_numbers

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return dfs(root, 0)
}

func dfs(TreeNode *TreeNode, preSum int) int {
	sum := preSum*10 + TreeNode.Val
	if TreeNode.Left == nil && TreeNode.Right == nil {
		return sum
	}
	if TreeNode.Left != nil && TreeNode.Right != nil {
		temp := sum
		sum = dfs(TreeNode.Left, temp)
		sum += dfs(TreeNode.Right, temp)
	} else if TreeNode.Left != nil {
		sum = dfs(TreeNode.Left, sum)
	} else {
		sum = dfs(TreeNode.Right, sum)
	}
	return sum
}
