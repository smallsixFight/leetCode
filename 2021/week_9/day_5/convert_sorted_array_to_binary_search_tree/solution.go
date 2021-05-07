package same_tree

import "github.com/leetCode/2021/common"

/*
108
来源：[leetCode](https://leetcode-cn.com/)
题目：[将有序数组转换为二叉搜索树](https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/)
标签：`树` `深度优先探索`

简单思路：先找到中位数作为根节点，然后将其左侧和右侧拆分为两个子数组，继续找中位数作为子树的根节点，以此类推。

官网运行结果记录
执行用时：0ms
内存消耗：3.4MB

*/

func sortedArrayToBST(nums []int) *common.TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	root := &common.TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root
}
