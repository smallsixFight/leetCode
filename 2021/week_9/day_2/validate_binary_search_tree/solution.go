package letter_combinations_of_a_phone_number

import "math"

/*
98
来源：[leetCode](https://leetcode-cn.com/)
题目：[验证二叉搜索树](https://leetcode-cn.com/problems/validate-binary-search-tree/)
标签：`树` `深度优先探索` `递归`

简单思路：这道题以前做过，使用的是 DFS 来做，先把当前结点跟左右两个自结点进行判断，成功后对左子树进行递归处理，并返回子树中的最大值和最小值，再进行二次判断，然后依次对右子树进行相同的处理。

这里重做就尝试了使用中序遍历来进行比较判断，这样就不用返回值，只需要判断遍历到的结点跟截至目前的最大值的比较，若小于则直接返回 false，遍历结束返回 true 即可。

官网运行结果记录
执行用时：8ms(isValidBST2)/8ms(isValidBST)
内存消耗：5.2MB(isValidBST2)/5.2MB(isValidBST)

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	minVal := math.MinInt64
	stack := make([]*TreeNode, 0)
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root, stack = stack[len(stack)-1], stack[:len(stack)-1]
		if root.Val <= minVal {
			return false
		}
		minVal = root.Val
		root = root.Right
	}
	return true
}

func isValidBST2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	isOk, _, _ := dfs(root)
	return isOk
}

func dfs(node *TreeNode) (bool, int, int) {
	minV, maxV := node.Val, node.Val
	if node.Left != nil && node.Left.Val >= node.Val {
		return false, 0, 0
	}
	if node.Right != nil && node.Right.Val <= node.Val {
		return false, 0, 0
	}
	if node.Left != nil {
		isOk, minVal, maxVal := dfs(node.Left)
		if !isOk || maxVal >= node.Val {
			return false, 0, 0
		}
		minV = minVal
	}
	if node.Right != nil {
		isOk, minVal, maxVal := dfs(node.Right)
		if !isOk || minVal <= node.Val {
			return false, 0, 0
		}
		maxV = maxVal
	}
	return true, minV, maxV
}
