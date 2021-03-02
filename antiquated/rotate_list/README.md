### 旋转链表（Rotate List）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/rotate-list/)

#### 题目
给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。

示例 1：
```
输入: 1->2->3->4->5->NULL, k = 2
输出: 4->5->1->2->3->NULL
解释:
向右旋转 1 步: 5->1->2->3->4->NULL
向右旋转 2 步: 4->5->1->2->3->NULL
```
示例 2：
```
输入: 0->1->2->NULL, k = 4
输出: 2->0->1->NULL
解释:
向右旋转 1 步: 2->0->1->NULL
向右旋转 2 步: 1->2->0->NULL
向右旋转 3 步: 0->1->2->NULL
向右旋转 4 步: 2->0->1->NULL
```

#### 思路一
由于移动的节点下一个位置不确定，所以使用两个 hash table 一个记录每个节点原来的值即可，一个记录移动后的值。

当移动后的值都确定后，就遍历修改链表的值即可。

为了避免 k 过大导致无意义的移动，需要找出有意义移动距离。

这题移动的是节点的 val，不需要改动原来的结构。

#### 思路二
通过思路一可以发现，其实只需要找到作为新的起点的节点，新起点的前一个节点则为新的尾节点，然后将原来的尾节点链接到原来的起点就可以了。

#### 伪代码
```
ROTATE-RIGHT-TWO(head, k)
    if head == nil
        return head
    l = 0
    for p != nil
        l ++
        p = p.next
    k %= l          // 有意义的移动数字
    if k == 0
        return head
    target = l +1 -k
    i = 1
    newNode = new ListNode()
    p = head
    while
        if i == target -1
            p1 = p
            p = p.next
            p1.next = nil
            continue
        elseif i == target
            newNode.next = p
        if p.next == nil
            p.next = head
            break
        p = p.next
    return newNode.next

ROTATE-RIGHT(head, k)
    if head == nil
        return head
    l = 0
    p = head
    record = new Map()
    newRecord = new Map()
    for p != nil
        l ++
        record[l] = p.val
        p = p.next
    k %= l          // 有意义的移动数字
    if k == 0
        return head
    for i = 1 to l
        idx = i +k
        if idx > l
            idx -= l
        newRecord[idx] = record[i]
    p = head
    idx = 1
    for p != nil
        p.val = newRecord[idx]
        idx ++
        p = p.next
    return head
```