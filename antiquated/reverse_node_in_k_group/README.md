### K 个一组翻转链表（Reverse Nodes in k-Group）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/reverse-nodes-in-k-group/)

#### 题目
给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。

k 是一个正整数，它的值小于或等于链表的长度。

如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

示例：
```
给你这个链表：1->2->3->4->5

当 k = 2 时，应当返回: 2->1->4->3->5

当 k = 3 时，应当返回: 3->2->1->4->5
```

说明：

- 你的算法只能使用常数的额外空间。
- **你不能只是单纯的改变节点内部的值**，而是需要实际进行节点交换。

#### 思路
首先实现链表的翻转。

可以每 k 个结点拆分成单独的链表，翻转后再组合起来；那么需要记录完成翻转后要链接的前一个结点的节点以及链接下一个翻转链表的尾节点，下一个拆分的起始位置。

#### 伪代码
```
REVERSE-K-GROUP-1(head, k)
    top = new ListNode()
    top.next = head
    pre, end = top, top
    while end != nil
        i = 0
        while i < k && end != nil
            end = end.next
        if end == nil
            break
        start = pre.next
        next = end.next
        end.next = nil
        pre.next = REVERSE-LIST(start)
        start.next = next
        pre = start
        end = pre
        i = 0
    return top.next
    
REVERSE-LIST(list)
    pre, next = nil, nil
    curr = list
    while curr != nil
        next = curr.next
        curr.next = pre
        pre = curr
        curr = next
    return pre
```