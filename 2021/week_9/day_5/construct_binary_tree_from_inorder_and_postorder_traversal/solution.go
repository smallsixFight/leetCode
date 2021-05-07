package construct_binary_tree_from_inorder_and_postorder_traversal

import "github.com/leetCode/2021/common"

/*
106
来源：[leetCode](https://leetcode-cn.com/)
题目：[从中序与后序遍历序列构造二叉树](https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/)
标签：`树` `深度优先探索` `数组`

简单思路：通过中序遍历和后序遍历的特性，可以知道，后序遍历的数组最后一位是根结点，而中序遍历中该元素左侧均为根节点左侧元素，且后序遍历数组中从头遍历的相同个数的结点也是左子树的结点。


官网运行结果记录
执行用时：8ms/12ms
内存消耗：4.2MB

*/

func buildTree(inorder []int, postorder []int) *common.TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	root := &common.TreeNode{Val: postorder[len(postorder)-1]}
	idx := 0
	for inorder[idx] != root.Val {
		idx++
	}
	root.Left = buildTree(inorder[:idx], postorder[:idx])
	root.Right = buildTree(inorder[idx+1:], postorder[idx:len(postorder)-1])
	return root
}
