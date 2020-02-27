package convent_sorted_array_to_binary_search_tree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printBinaryTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.Val, ", ")
	printBinaryTree(root.Left)
	printBinaryTree(root.Right)
}

func SortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	node := &TreeNode{Val: nums[len(nums)/2]}
	node.Left = SortedArrayToBST(nums[0 : len(nums)/2])
	node.Right = SortedArrayToBST(nums[len(nums)/2+1:])
	return node
}
