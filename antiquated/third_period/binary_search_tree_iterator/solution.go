package binary_search_tree_iterator

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	iterator := BSTIterator{stack: make([]*TreeNode, 0)}
	for root != nil {
		iterator.stack = append(iterator.stack, root)
		root = root.Left
	}
	return iterator
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	node := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	if node.Right != nil {
		this.stack = append(this.stack, node.Right)
		p := node.Right.Left
		for p != nil {
			this.stack = append(this.stack, p)
			p = p.Left
		}
	}
	return node.Val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}
