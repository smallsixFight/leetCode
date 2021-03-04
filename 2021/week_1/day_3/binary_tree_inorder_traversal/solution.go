package binary_tree_inorder_traversal

/*
来源：[leetCode](https://leetcode-cn.com/)
题目：[二叉树的中序遍历](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/)
标签：`栈` `树` `哈希表`

简单思路：中序遍历就是先读取最左孩子节点的值，然后读取当前节点的值，最后返回右孩子节点的值。用递归很容易实现，所以
描述下迭代（DFS）的逻辑。
首先，遍历的顺序是先进行左结点进行不断遍历直到叶子结点，而遍历的结点后续回溯时才能取值，所以需要用栈存起来。
每取出一个栈顶结点后，将 Val 加入结果集，然后对其的右结点进行判断，如果右结点存在，则入栈并遍历其左结点，直到栈
变空。此外，栈顶结点被取出时，其左子树是遍历完成的，不能够再次遍历，因此需要一个标识用作判断。
有两种情况才能从栈取值：左结点遍历到叶子结点或栈顶结点没有右结点。

官网运行结果记录
执行用时：0ms
内存消耗：2MB

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	p := root
	// p 指向未遍历的结点。stack 不为空时，则可以取值并获取不为空的右结点作为遍历的结点。
	for p != nil || len(stack) > 0 {
		var isPop bool // 是否取出栈顶的结点
		if p == nil {
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			isPop = true
		}
		if !isPop && p.Left != nil { // 不是从栈顶取结点且左结点不为空，继续遍历
			stack = append(stack, p)
			p = p.Left
			continue
		}
		res = append(res, p.Val)
		p = p.Right // 右结点为空时，需要从栈取结点，不为空则遍历其左结点
	}
	return res
}

// 递归
func inorderTraversal2(root *TreeNode) []int {
	res := make([]int, 0)
	if root != nil {
		res = append(res, inorderTraversal2(root.Left)...)
		res = append(res, root.Val)
		res = append(res, inorderTraversal2(root.Right)...)
	}
	return res

}
