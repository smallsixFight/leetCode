###  排序链表（Sort）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/sort-list/)

#### 题目
在`O(n log n)`时间复杂度和常数级空间复杂度下，对链表进行排序。

示例 1：
```
输入: 4->2->1->3
输出: 1->2->3->4
```
示例 2：
```
输入: -1->5->3->4->0
输出: -1->0->3->4->5
```

#### 思路一
看到复杂度要求，首先想到的就是快速排序、归并排序和堆排序，由于这是链表排序，就不先考虑空间复杂度，不然的话就只有堆排序可选了。

我先考虑一下能否将快速排序做一下调整，可以想到，链表是指针，所以不需要额外的空间，有尝试的可能性。

以链表的第一个节点作为分割点，将小于该节点的值的节点移动到左边，大于的移动到右边，由于是以第一个节点为分割点，大于的节点并不需要操作。

#### 思路二
在把快排的版本提交后，可能我写的烂，没有去找中间分割点，导致超出时间限制（比我前一道题的插入排序还扎心），所以思路转向了具有稳定性的归并排序。

因为是链表，可是进行 cut 和 join，所以可以只需要常数的空间，不过需要考究就是 cut 和 join 的操作，以及找到分割的中间位置。

中间位置可以通过快慢指针的方式来寻找。

#### 伪代码
```
SORT-LIST-WITH-MERGE(head)
    if head == nil || head.next == nil
        return head
    p1, p2, head, head
    for p2 != nil && p2.next != nil
        p1 = p1.next
        p2 = p2.next.next
    l1 = SORT-LIST-WITH-MERGE(p1.next)
    p1.next = nil                   // 断开
    l2 = SORT-LIST-WITH-MERGE(head)
    if l1 == nil
        return l2
    elseif l2 == nil
        return l1
    top = new ListNode()
    top.next = l1
    p3 = top
    for l2 != nil
        if p3.next == nil
            p3.next = l2
            break
        elseif l2.val < p3.next.val
            temp = l2.next
            l2.next = p3.next
            p3.next = l2
            l2 = temp
        p3 = p3.next
    return top.next

SORT-LIST-WITH-QUICK(head)
    top = new ListNode()
    top.next = head
    SORT(top, nil)
    return top.next
    
SORT(top, end)
    if top.next == nil || top.next == end
        return
    n = top.next
    e1 = top.next
    for n.next != end
        if n.next.val < top.next.val
            temp = n.next
            n.next = n.next.next
            temp.next = top.next
            top.next = temp
        else
            n = n.next
    SORT(top, e1)
    SORT(e1, nil)
```