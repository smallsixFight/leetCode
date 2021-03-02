package convert_sorted_list_to_binary_search_tree

import "strconv"

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getTreeArr(tree *TreeNode) []string {
	if tree == nil {
		return nil
	}
	res := make([]string, 0)
	stack := []*TreeNode{tree}
	for len(stack) > 0 {
		l := len(stack)
		for i := 0; i < l; i++ {
			node := stack[i]
			if node != nil {
				res = append(res, strconv.FormatInt(int64(node.Val), 10))
				stack = append(stack, node.Left)
				stack = append(stack, node.Right)
			} else {
				res = append(res, "null")
			}
		}
		stack = stack[l:]
	}
	return res
}

func SortedListToBFT(head *ListNode) *TreeNode {
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

func dfs(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	mid := len(arr) / 2
	node := &TreeNode{Val: arr[mid]}
	node.Left = dfs(arr[:mid])
	node.Right = dfs(arr[mid+1:])
	return node
}
