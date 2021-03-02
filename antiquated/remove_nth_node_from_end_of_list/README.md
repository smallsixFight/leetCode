### 删除链表的倒数第 N 个节点（Remove Nth Node From End of List）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/)

#### 题目
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：
```
给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
```

**说明：**  
给定的 _n_ 保证是有效的。

**进阶：**  
你能尝试使用一趟扫描实现吗？

#### 思路
这题我直接写了两个一次遍历的方法，一个是使用 hash 表，一个使用双指针法。

遍历并将新建一个指针存储在指针数组里，然后根据 n 获取其前一个节点，调整前个节点的 Next。

通过前一种方法可以发现，我们仅需要与最后一个结点相差距离为 n 的另一个节点就好了；那么使用两个指针和一个计数值 count。  
当`count == n+1`时，指向需要修改 Next 结点的指针为 head；当`count > n +1`时，`pre=pre.Next`。

#### 伪代码
```
REMOVE-NTH-FROM-END-2(head, n)
    if head == nil
        return nil
    pre, right, count = nil, head, 1
    for right.Next != nil
        if count < n +1
            count += 1
        elseif pre == nil
            pre = head
        else
            pre = pre.Next
        right = right.Next
    if p1.Next != nil
        p1.Next = p1.Next.Next
    else
        head = head.Next
    return head   

REMOVE-NTH-FROM-END(head, n)
    if head == nil
        return nil
    arr = new listNode Array()
    p = head.Next
    for p != nil
        arr.push(p)
    idx = arr.length -1 -n
    if idx < 0
        head.next = head.next.next
    else
        arr[idx].next = arr[idx].next.next
    return head
```