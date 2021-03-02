### 合并两个有序链表（Merge Two Sorted Lists）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/merge-two-sorted-lists/)

#### 题目
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：
```
输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
```

#### 思路
这题可以先创建一个 head 和 pre 结点，head 的 Next 指向返回链表的第一个结点，pre 指向最后一个加入返回链表的结点。  
另外，本题可以不创建新的结点，而是最后一个加入的结点的 Next 指向就好。

#### 伪代码
```
MERGE-TWO-SORTED-LISTS(l1, l2)
    head = new ListNode{}
    pre = new ListNode{}
    pre = head
    while l1 != nil || l2 != nil
        if l1 == nil
            pre.Next = l2
            l2 = l2.Next
        elseif l2 == nil
            pre.Next = l1
            l1 = l1.Next
        else
            if l1.Val < l2.Val
                pre.Next = l1
                l1 = l1.Next
            else
                pre.Next = l2
                l2 = l2.Next
        pre = pre.Next
    return head.Next
```