package next_greater_node_in_linked_list

/*
1019
来源：[leetCode](https://leetcode-cn.com/)
题目：[链表中的下一个更大节点](https://leetcode-cn.com/problems/next-greater-node-in-linked-list/)
标签：`栈` `链表`

简单思路：这道题目与 `下一个更大元素 I` 类似，只是将数组改成了链表，但其实做法一模一样，可以用一个单调栈来处理。

官网运行结果记录
执行用时：68ms(68.8%)
内存消耗：7.7MB(14.4%)

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func nextLargerNodes(head *ListNode) []int {
	res, stack := make([]int, 0), make([][2]int, 0)
	for head != nil {
		for len(stack) > 0 && head.Val > stack[len(stack)-1][0] {
			res[stack[len(stack)-1][1]] = head.Val
			stack = stack[:len(stack)-1]
		}
		res = append(res, 0)
		stack = append(stack, [2]int{head.Val, len(res) - 1})
		head = head.Next
	}

	return res
}
