package binary_search_three_iterator

/*
来源：[leetCode](https://leetcode-cn.com/)
题目：[二叉搜索树迭代器](https://leetcode-cn.com/problems/binary-search-tree-iterator/)
标签：`栈` `树` `设计`

简单思路：首先需要重新了解下什么叫做 `二叉搜索树` 才行。二叉搜索树又叫二叉查找树，左子树上的结点均小于根结点，右子树上的结点均大于根结点。那么可以对树进行中序遍历，使用一个队列来存储，每次取值取队头就可以了。
如果要按照标签的 `栈` 来实现，那么就简单的异构一下中序遍历，将 `左-根-右` 改为 `右-根-左` 遍历入栈。然而这个思路不符合题目空间复杂度为 O(h) 的要求。

我没有想出能够在 O(h) 空间范围内将 next 维持在 O(1) 的方法。在看了官方题解的思路后，才知道通过均摊来将 next 的平均时间复杂度达到 O(1)，但 next 的最坏时间复杂度也就是 O(n)。

官网运行结果记录
执行用时：28ms		24ms/28ms（均摊思路）
内存消耗：8.9MB/9MB		9.1MB（均摊思路）

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	stack []int
}

func Constructor(root *TreeNode) BSTIterator {
	iterator := BSTIterator{}
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Right
		}
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		iterator.stack = append(iterator.stack, p.Val)
		if p.Left != nil {
			root = p.Left
		}
	}
	return iterator
}

func (this *BSTIterator) Next() int {
	v := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	return v
}

func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}

// BSTIterator2 均摊写法
type BSTIterator2 struct {
	stack []*TreeNode
}

func (this *BSTIterator2) inorder(root *TreeNode) {
	for root != nil {
		this.stack = append(this.stack, root)
		root = root.Left
	}
}

func Constructor2(root *TreeNode) BSTIterator2 {
	iterator := BSTIterator2{}
	iterator.inorder(root)
	return iterator
}

func (this *BSTIterator2) Next() int {
	node := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	if node.Right != nil {
		this.inorder(node.Right)
	}
	return node.Val
}

func (this *BSTIterator2) HasNext() bool {
	return len(this.stack) > 0
}
