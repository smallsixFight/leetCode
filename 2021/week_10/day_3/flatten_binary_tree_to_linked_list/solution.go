package leaf_similar_trees

import (
	"github.com/leetCode/2021/common"
)

/*
114
来源：[leetCode](https://leetcode-cn.com/)
题目：[二叉树展开为链表](https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/)
标签：`树` `深度优先探索`

简单思路：这道题在 O(n) 的复杂度可以很容易想到题解，通过前序遍历获取对应结点顺序，然后修改结点的左右子结点指针。

官网运行结果记录
执行用时：4ms(flatten)
内存消耗：2.9MB(flatten)

*/

func flatten(root *common.TreeNode) {
	list, stack := make([]*common.TreeNode, 0), make([]*common.TreeNode, 0)
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			list = append(list, node)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node = node.Right
	}
	for i := 1; i < len(list); i++ {
		prev, cur := list[i-1], list[i]
		prev.Left, prev.Right = nil, cur
	}
}
