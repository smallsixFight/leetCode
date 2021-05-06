package construct_binary_tree_from_preorder_and_inorder_traversal

import "github.com/leetCode/2021/common"

/*
105
来源：[leetCode](https://leetcode-cn.com/)
题目：[从前序与中序遍历序列构造二叉树](https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/)
标签：`树` `深度优先探索` `数组`

简单思路：根据前序遍历和中序遍历的规律，可以知道，前序遍历的第一个元素为根节点，且后续分别为其左子树的所有结点和右子树所有结点两部分；
而中序遍历的得到的数组中，根节点前面的元素为其左子树结点，后面的元素为右子树结点。那么可以通过递归来拆分构建二叉树。

官网运行结果记录
执行用时：4ms
内存消耗：4.2MB

*/

func buildTree(preorder []int, inorder []int) *common.TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &common.TreeNode{Val: preorder[0]}
	idx := 0
	for ; idx < len(inorder); idx++ {
		if inorder[idx] == preorder[0] {
			break
		}
	}
	root.Left = buildTree(preorder[1:idx+1], inorder[:idx])
	root.Right = buildTree(preorder[idx+1:], inorder[idx+1:])
	return root
}
