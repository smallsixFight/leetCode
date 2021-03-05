package binary_tree_zigzag_level_order_traversal

/*
来源：[leetCode](https://leetcode-cn.com/)
题目：[二叉树的锯齿形层序遍历](https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/)
标签：`栈` `树` `广度优先探索`

简单思路：先想一下每层都是从左到右遍历怎么做：使用两个队列，一个队列保存目前需要读取的结点，
一个队列保存下一层要读取的结点。正在读取的结点，每次取值后依次将其左结点和右结点
放入下层队列中，以此轮换。在 go 中，两个队列可以用一个 slice 代替，因为每次读取的
长度大小是确定的，事先确定好长度就行。
通过上面的方法简单扩展一下，如果添加一个 flag 来表示这次是从左到右读取，还是从右
到左读取，那么就完成了。
题目的有个 `栈` 的标签，用栈的话取值的顺序就乱了，但前面的方法可以通过 flag 来判断
从哪边先读取，那么通过 flag 判断这层是先 push 左结点还是右结点也就解决了。

官网运行结果记录
执行用时：0ms
内存消耗：2.5MB

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	stack := []*TreeNode{root}
	var isReverse bool
	for len(stack) > 0 {
		l := len(stack)
		temp := make([]int, 0)
		for i := l - 1; i >= 0; i-- { // 模拟出栈
			n := stack[i]
			temp = append(temp, n.Val)
			if !isReverse {
				if n.Left != nil {
					stack = append(stack, n.Left)
				}
				if n.Right != nil {
					stack = append(stack, n.Right)
				}
			} else {
				if n.Right != nil {
					stack = append(stack, n.Right)
				}
				if n.Left != nil {
					stack = append(stack, n.Left)
				}
			}
		}
		isReverse = !isReverse
		stack = stack[l:]
		res = append(res, temp)
	}
	return res
}
