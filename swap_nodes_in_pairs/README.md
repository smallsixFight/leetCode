### 两两交换链表中的节点（Swap Nodes in Pairs）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/swap-nodes-in-pairs/)

#### 题目
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

**你不能只是单纯的改变节点内部的值**，而是需要实际的进行节点交换。

示例：
```
给定 1->2->3->4, 你应该返回 2->1->4->3.
```

#### 思路
这题不仅是要处理两个交换结点的 next 值，还需要处理第一个节点的 next 值，否则第二个节点就会消失。  

由于提供的 struct 没有 p 指针，所以我们需要额外用一个结点来记录。

由于第一个结点可能第二个交换，所以需要一个 top 指向链表头部。

#### 伪代码
```
SWAP-PAIRS(head)
    if head == nil
        return head
    top = new ListNode()
    top.next = head
    pre = top
    for pre.next != nil
        if pre.next.next != nil
            p1, p2 = pre.next, pre.next.next
            pre.next = p2
            p1.next = p2.next
            p2.next = p1
            pre = pre.next.next
        else
            break
    return top.next
```