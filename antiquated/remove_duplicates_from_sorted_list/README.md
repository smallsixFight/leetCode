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
`DeleteDuplicates2()`是我的版本，我实在看不出两个版本的差异在哪里，本地测试两个结果又一样，很懵比就是了。

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