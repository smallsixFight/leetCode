package unique_binary_search_trees_II

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PreOrderPrint(root *TreeNode) []string {
	res := make([]string, 0)
	if root != nil {
		res = append(res, strconv.FormatInt(int64(root.Val), 10))
		res = append(res, PreOrderPrint(root.Left)...)
		res = append(res, PreOrderPrint(root.Right)...)
	} else {
		res = append(res, "null")
	}
	return res
}

func GenerateTrees(n int) []*TreeNode {
	if n < 1 {
		return nil
	}
	res := make([]*TreeNode, 0)
	res = append(res, &TreeNode{Val: 1})
	createNewTree := func(tree *TreeNode, n int) {
		t := tree
		var pre *TreeNode
		for t != nil {
			newChild := &TreeNode{}
			if pre != nil {
				temp := t
				pre.Right = nil
				newTree := &TreeNode{}
				copyTree(tree, newTree)
				copyTree(t, newChild)
				node := &TreeNode{Val: n, Left: newChild}
				p := newTree
				for p.Right != nil {
					p = p.Right
				}
				p.Right = node
				res = append(res, newTree)
				pre.Right = temp
			} else {
				copyTree(t, newChild)
				node := &TreeNode{Val: n, Left: newChild}
				res = append(res, node)
			}
			pre = t
			t = t.Right
		}
		pre.Right = &TreeNode{Val: n}
	}
	for i := 2; i <= n; i++ {
		l := len(res)
		for k := 0; k < l; k++ {
			createNewTree(res[k], i)
		}
	}
	return res
}

func copyTree(oldTree, newTree *TreeNode) {
	if oldTree != nil {
		newTree.Val = oldTree.Val
	} else {
		newTree = nil
		return
	}
	if oldTree.Left != nil {
		newTree.Left = &TreeNode{}
		copyTree(oldTree.Left, newTree.Left)
	}
	if oldTree.Right != nil {
		newTree.Right = &TreeNode{}
		copyTree(oldTree.Right, newTree.Right)
	}
}
