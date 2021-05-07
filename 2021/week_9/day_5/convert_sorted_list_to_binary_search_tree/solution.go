package same_tree

import "github.com/leetCode/2021/common"

/*
109
来源：[leetCode](https://leetcode-cn.com/)
题目：[有序链表转换二叉搜索树](https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree/)
标签：`深度优先探索` `链表`

简单思路：这道题跟前面的 `将有序数组转换为二叉搜索树` 类似，区别就在于如何找到链表的中间位置，可以使用快慢指针法来解决这个问题。

官网运行结果记录
执行用时：12ms
内存消耗：6.4MB

*/

func sortedListToBST(head *common.ListNode) *common.TreeNode {
	if head == nil {
		return nil
	}
	list := &common.ListNode{}
	p, f := list, head
	for f != nil && f.Next != nil {
		p.Next = &common.ListNode{Val: head.Val}
		p = p.Next
		head = head.Next
		f = f.Next.Next
	}
	root := &common.TreeNode{Val: head.Val}
	root.Left = sortedListToBST(list.Next)
	root.Right = sortedListToBST(head.Next)
	return root
}

// 转换成数组处理
func sortedListToBST2(head *common.ListNode) *common.TreeNode {
	if head == nil {
		return nil
	}
	arr := make([]int, 0)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	return dfs(arr)
}

func dfs(arr []int) *common.TreeNode {
	if len(arr) == 0 {
		return nil
	}
	mid := len(arr) / 2
	node := &common.TreeNode{Val: arr[mid]}
	node.Left = dfs(arr[:mid])
	node.Right = dfs(arr[mid+1:])
	return node
}
