### 两数相加（Add Two Numbers）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/add-two-numbers/)

#### 题目
给出两个`非空`的链表用来表示两个非负的整数。其中，它们各自的位数是按照`逆序`的方式存储的，并且它们的每个节点只能存储`一位`数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：
```
输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
```

#### 思路
这道题只是单纯考链表的操作，只需要考虑下进位问题。

#### 伪代码
ADD-TWO-NUMBERS(l1, l2)
    head = new ListNode()
    p = head
    carry = 0
    for l1 != nil || l2 != nil || carry != 0
        temp = 0 + carry
        if l1 != nil
           temp += l1.val
           l1 = l1.next
        if l2 != nil
            temp != l2.val
            l2 = l2.next
        if temp > 9
            carry = 1
        else 
            carry = 0
        node = new ListNode(temp %10)
        p.next = node
        p = p.next 
    return head.next
