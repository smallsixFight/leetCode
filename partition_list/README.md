### 分隔链表（Partition List）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/partition-list/)

#### 题目
给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。

你应当保留两个分区中每个节点的初始相对位置。

示例：
```
输入: head = 1->4->3->2->5->2, x = 3
输出: 1->2->2->4->3->5
```

#### 思路
题目的`保留每个节点的初始相对位置`算是一个提示，要保证这一点的话，只需要修改 Next 就可以了。

大致思路如下：

- 使用一个指针`top`作为链表头；
- 使用一个指针`pre`记录第一个遇到大于等于 x 的前一个节点位置；
- 使用一个指针`p`来进行遍历；
- 当`p.next.val`上的值小于 x 并且`p>pre`时，`pre.next = p.next`、`p.next = p.next.next`，交换结束后`pre`向后移动一位，然后`p`继续向后遍历。

上面最后一步具体看伪代码，当中还需要使用一些临时变量。


#### 伪代码
```
PARTITION(head, x)
    if head == nil
        return head
    top = new ListNode()
    top.next = head
    pre = top
    p = top
    while p.next != nil
        if p.next.val >= x && pre == top
            pre = p
        elseif p.next.val < x
            if pre != p
                temp = pre.next
                pre.next = p.next
                p.next = p.next.next
                pre.next.next = temp
                pre = pre.next
                continue            // 因为 next 上额值做了交换，不能向后移动
            pre = pre.next
        p = p.next
    return top.next
```