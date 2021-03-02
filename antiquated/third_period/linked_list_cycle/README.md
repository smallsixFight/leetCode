### 环形链表（Linked List Cycle）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/linked-list-cycle/)

#### 题目
给定一个链表，判断链表中是否有环。

为了表示给定链表中的环，我们使用整数`pos`来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果`pos`是 -1，则在该链表中没有环。

示例 1：
```
输入：head = [3,2,0,-4], pos = 1
输出：true
解释：链表中有一个环，其尾部连接到第二个节点。
```
![circularlinkedlist.png](https://i.loli.net/2020/04/22/Z8CIKDrwVGihBHR.png)

示例 2：
```
输入：head = [1,2], pos = 0
输出：true
解释：链表中有一个环，其尾部连接到第一个节点。
```
![circularlinkedlist_test2.png](https://i.loli.net/2020/04/22/i129lesPDtnRc4F.png)

示例 3：
```
输入：head = [1], pos = -1
输出：false
解释：链表中没有环。
```
![circularlinkedlist_test3.png](https://i.loli.net/2020/04/22/DnIAzXW2RcdHaBb.png)

进阶：

你能用 O(1)（即，常量）内存解决此问题吗？

#### 思路
在这一期的中等题遇到过这道题的扩展，这道题只需要判断是否存在环就行，相对而言少了一个步骤。

#### 伪代码
HAS-CYCLE(head)
    if head == nil
        return head
    p1, p2 = head, head.next
    for p1 != p2
        if p2 != nil && p2.next != nil
            p2 = p2.next.next
        else
            return false
        p1 = p1.next
    return true