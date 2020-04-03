### 反转链表 II（Reverse Linked List II）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/sort-colors/)

#### 题目
反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。

说明：  
1 ≤ m ≤ n ≤ 链表长度。

示例：
```
输入: 1->2->3->4->5->NULL, m = 2, n = 4
输出: 1->4->3->2->5->NULL
```

#### 思路
把 m-n 的链表拉出来然后做一个反转，再接回原来的链表中。

思路很简单，但难度是在于对链表的处理，大致整理如下：

- 使用 top 指向链表头，因为第一个元素也可能被反转；
- 使用 pre 指向开始位置的前一个节点，用于恢复链表；
- 使用 s 指向开始位置，用于反转后更新 next 值；
- 使用 p 用于遍历数组；
- 遍历反转子链表，遍历结束更新 pre 和 s 的 next 值。

#### 题外话
这道题要完整的在一遍循环遍历里面处理好每个节点的关系让我脑壳疼，一开始写的伪代码是错误的，直接死循环了，后面又重新画了草图，重新思考逻辑才通过。

另外，我本想着看看别人提交的代码看看能不能优化下代码，然而看到的几个例子都用了三个 for 或者两个 for，因此暂时作罢（只看了 golang 的）。

#### 伪代码
```
REVERSE-BETWEEN(head, m, n)
    if head == nil || head.next == nil || m >= n
        return head
    top.next = head
    pre, sp = top, top
    s, p = head, head
    count = 1
    while p != nil
        if count < m
            s = s.next
            sp = sp.next
            pre = pre.next
        elseif count == m
            pre = pre.next
        elseif count <= n
            temp = p.next
            p.next = pre
            pre = p
            if count == n
                sp.next = p
                s.next = temp
                break
            p = temp
            count ++
            continue
        count ++
        p = p.next
    return top.next
```