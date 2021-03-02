### 删除排序链表中的重复元素 II（Remove Duplicates from Sorted List II）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/)

#### 题目
给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。

示例 1：
```
输入: 1->2->3->3->4->4->5
输出: 1->2->5
```
示例 2：
```
输入: 1->1->1->2->3
输出: 2->3
```

#### 思路
本题大致思路如下：

- 使用`pre`来记录前一个值，初始化为`head.val -1`；
- 使用`top`记录返回的起始位置，初始化为`nil`；
- 使用`p1`来指向当前已经最后一个不重复数字的节点，初始化为`nil`；
- 使用`p2`来迭代遍历，初始化指向`head`。

当`p2.val != pre`时：

- 当`(p2.next != nil && p2.val != p2.next.val) || p2.next == nil`时，如果`top == nil`，`top = p2; p1 = p2`，否则`p1.next = p2; p1 = p1.next`;；
- `pre = p2.val`。

#### 伪代码
```
DELETE-DUPLICATES(head)
    if head == nil || head.next == nil
        return head
    top, p1, p2 = nil, nil, head
    pre := head.val -1
    for p2 != nil
        if p2.val != pre
            if (p2.next != nil && p2.val != p2.next.val) || p2.next == nil
                if top == nil
                    top = p2
                    p1 = p2
                else
                    p1.next = p2
                    p1 = p1.next
            pre = p2.val
        p2 = p2.next
        if p1 != nil        // 在不确定是否存在下一个节点时，置 0
            p1.next = nil
    return top
```