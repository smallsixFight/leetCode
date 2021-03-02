### 合并 K 个排序链表（Merge k Sorted Lists）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/merge-k-sorted-lists/)

#### 题目
合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。

示例：
```
输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6
```

#### 思路
这道题可以使用分治法来做， 将链表数组分到一到两个链表后再进行合并，返回合并后的链表，以此递归。

每两个链表合并的复杂度为`O(N)`，k 个链表需要进行`k-1`次合并，所以时间复杂度为`O(log(K))`。

#### 题外话
`MergeKLists2()`是逐一合并两个链表的方法，想对于分治法，只是改动了一点细微的地方，但是效率确实天差地别。

#### 伪代码
```
MERGE-K-LISTS(lists)
    if lists.length == 1
        return lists[0]
    l1, l2 = lists[0], lists[1]
    if lists.length > 2
        mid = lists.length /2 +1
        l1 = MERGE-K-LISTS(lists[0:mid])
        l2 = MERGE-K-LISTS(lists[mid:])
    top = new ListNode()
    p = top
    for l1 != nil || l2 != nil
        if l1 != nil && l2 != nil
            if l1.val < l2.val
                p.next = l1
                l1 = l1.next
            else
                p.next = l2
                l2 = l2.next
        elseif l1 != nil
            p.next = l1
            l1 = l1.next
        else
            p.next = l2
            l2 = l2.next
        p = p.next
    return top.next
```