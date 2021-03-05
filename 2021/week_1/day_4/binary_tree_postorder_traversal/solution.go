package binary_tree_postorder_traversal

/*
来源：[leetCode](https://leetcode-cn.com/)
题目：[二叉树的后序遍历](https://leetcode-cn.com/problems/binary-tree-postorder-traversal/)
标签：`栈` `树`

简单思路：后序遍历即先处理其左右结点，然后再处理当前结点。
把根结点入栈，指针移动到它的左结点，如果没有左结点。如果指针指向的为 nil，则从栈
取出一个新的结点。从栈取出的结点需要先判断是否有右结点，有则将当前结点再次入栈，指针指向右结点，重复开始的步骤。
当从栈取出的结点没有右结点或右结点已经遍历完成，则将 Val 加入结果集，并继续从栈取结点。
这时有跟中序遍历一样的问题，就是右结点会被重复遍历，陷入死循环，这时也一样需要一个 flag 来处理。
当从栈顶取一个结点，需要先判断并处理它的右结点，之后再处理当前结点，意味当前结点可能在第二次回溯到
的时候才会被处理，这与中序遍历不同，所以需要使用一个 hashTable 来记录：
栈顶结点被出来时，如果该结点已被记录或没有右结点，则直接取值和 push，否则，进行记录，并将指针指向右结点。

根据题解扩展思路：
当一个结点被取值时，意味着以当前结点为根结点的整棵树都遍历完成了，那么它的父结点就不用对其进行判断遍历。
当从栈取出的结点的右结点存在但已被取值时，则不需要对右结点做处理。

官网运行结果记录
执行用时：0ms
内存消耗：2MB

*/

// todo: Morris 遍历

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	record := make(map[*TreeNode]bool)
	stack := make([]*TreeNode, 0)
	p := root
	for p != nil || len(stack) > 0 {
		if p != nil {
			stack = append(stack, p)
			p = p.Left
			continue
		}
		p = stack[len(stack)-1]
		if record[p] || p.Right == nil {
			res = append(res, p.Val)
			stack = stack[:len(stack)-1]
			p = nil
		} else {
			record[p] = true
			p = p.Right
		}

	}
	return res
}

func postorderTraversal2(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	res := make([]int, 0)
	pre := &TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		if root.Right == nil || root.Right == pre {
			res = append(res, root.Val)
			stack = stack[:len(stack)-1]
			pre = root
			root = nil
		} else {
			root = root.Right
		}
	}
	return res
}
