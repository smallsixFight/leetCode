### 重排链表（Reorder List）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/reorder-list/)

#### 题目
给定一个单链表 L：L0→L1→…→Ln-1→Ln ，  
将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例 1：
```
给定链表 1->2->3->4, 重新排列为 1->4->2->3.
```
示例 2：
```
给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.
```

#### 思路
使用一个数组保存每个节点的指针位置，然后用双指针去修改 next 值。

#### 伪代码
```
REORDER-LIST(head)
    if head == nil
        return head
    arr = new Array()
    while head != nil
        arr.push(head)
        head = head.next
    p1, p2 = 0, arr.length -1
    while p1 < p2
        arr[p1].next = arr[p2]
        arr[p2].next = arr[p1+1]
        p1 ++
        p2 --
    arr[p1].next = nil
```