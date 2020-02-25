### 删除排序链表中的重复元素（Remove Duplicates from Sorted List）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/)

#### 题目
给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

示例 1：
```
输入: 1->1->2
输出: 1->2
```
示例 2：
```
输入: 1->1->2->3->3
输出: 1->2->3
```

#### 思路
该题使用单指针的方式来解决，当 p1 指向的位置上的值与 Next 上的值相等时，调整`p1.Next = p1.Next.Next`；不相等时则 p1 向后移动。

#### 题外话
不知道为什么 LeetCode 提交后没有通过测试案例，我的代码与官方的 Java 几乎相同，而且本地测试没问题，不管他了。

#### 伪代码
```
DELETE-DUPLICATES(head)
    if head == nil || head.Next == nil
        return head
    p1 = head.Next
    for p1.Next != nil
        if p1.Val == p1.Next.Val
            p1.Next = p1.Next.Next
        else
            p1 = p1.Next
    return head
```