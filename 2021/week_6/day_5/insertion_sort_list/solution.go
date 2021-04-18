package sort_colors

/*
147
来源：[leetCode](https://leetcode-cn.com/)
题目：[对链表进行插入排序](https://leetcode-cn.com/problems/insertion-sort-list/)
标签：`排序` `链表`

简单思路：对我来说，这道题的难点还在在于指针的处理。大致的处理逻辑如下：

- 若 head 为空结点或之后一个结点，直接返回 head；
- 使用一个新结点 top，top.next -> head，保证遍历结束后不会丢失 head 最开始的为止，并且能交换 head 与后续结点；
- 使用两个指针 p1、p2 分别指向 head、head.next，然后开始遍历比较 p1、p2 指向的值；
- 若 p2.Val < p1.Val，p1.next 指向 p2.next，从头遍历 top，查找适合插入 p2 的位置，插入后要调整对应的 next 指针值；
- 遍历结束后，返回 top.next。


官网运行结果记录
执行用时：4ms
内存消耗：3.2MB

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	top := &ListNode{Next: head}
	p1, p2 := head, head.Next
	for p2 != nil {
		if p2.Val < p1.Val {
			temp := p2.Next
			p1.Next = p2.Next
			p3 := top
			for p3.Next.Val < p2.Val {
				p3 = p3.Next
			}
			p2.Next = p3.Next
			p3.Next = p2
			p2 = temp
		} else {
			p1 = p1.Next
			p2 = p2.Next
		}
	}
	return top.Next
}
