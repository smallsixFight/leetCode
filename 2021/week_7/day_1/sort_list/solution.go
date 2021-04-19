package sort_list

/*
148
来源：[leetCode](https://leetcode-cn.com/)
题目：[排序链表](https://leetcode-cn.com/problems/sort-list/)
标签：`排序` `链表`

简单思路：先尝试插入排序，如 `sortList()` 所示，时间消耗很大，根据题目时间复杂度为 O(nlog n) 的提示，可以想到归并排序和堆排序，堆排序好像不适合，归并排序的空间复杂度确实 O(n)。
但是，归并排序的空间复杂度是按照数据来考虑的，根据链表的特性，归并排序的空间复杂度是常数范围的。归并排序的处理方式大致如下：

- 使用双指针法来寻找链表的中间结点，进行`截断`，通过递归进行再处理，直到再处理的部分只有一个结点或没有结点可处理；
- 然后对截断后链表进行插入排序合并为一个链表；
- 最后不断回溯处理直到第一层返回结果。

官网运行结果记录
执行用时：308ms(sortList)/24ms(sortList2)/32ms(sortList2)
内存消耗：7.1MB(sortList)/7.1MB(sortList2)

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	prev, p1, p2 := &ListNode{}, head, head
	for p2 != nil && p2.Next != nil {
		prev = p1
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	prev.Next = nil // 截断
	l1 := sortList2(p1)
	l2 := sortList2(head)
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	top := &ListNode{Next: l1}
	p := top
	for l2 != nil {
		if p.Next == nil {
			p.Next = l2
			break
		} else if l2.Val < p.Next.Val {
			temp := l2.Next
			l2.Next = p.Next
			p.Next = l2
			l2 = temp
		}
		p = p.Next
	}
	return top.Next
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	top := &ListNode{Next: head}
	p := head
	for p.Next != nil {
		if p.Val > p.Next.Val {
			temp := p.Next
			p.Next = p.Next.Next
			p2 := top
			for p2.Next.Val < temp.Val {
				p2 = p2.Next
			}
			temp.Next = p2.Next
			p2.Next = temp
		} else {
			p = p.Next
		}
	}
	return top.Next
}
