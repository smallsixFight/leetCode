package range_sum_of_bst

/*
938
来源：[leetCode](https://leetcode-cn.com/)
题目：[二叉搜索树的范围和](https://leetcode-cn.com/problems/range-sum-of-bst/)
标签：`排序`

简单思路：进行中遍历，比较汇总。

官网运行结果记录
执行用时：100ms(96%)
内存消耗：8.6MB(59%)

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	res := 0
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		if node.Left != nil {
			stack = append(stack, node.Left)
			node.Left = nil // 不再重复遍历左结点
		} else {
			if node.Val >= low && node.Val <= high {
				res += node.Val
			}
			stack = stack[:len(stack)-1]
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
	}
	return res
}
