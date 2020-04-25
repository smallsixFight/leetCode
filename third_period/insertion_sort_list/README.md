### 对链表进行插入排序（Insertion Sort List）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/insertion-sort-list/)

#### 题目
对链表进行插入排序。

![Insertion-sort-example-300px.gif](https://i.loli.net/2020/04/25/VqWvlOgJcnBjZha.gif)

插入排序的动画演示如上。从第一个元素开始，该链表可以被认为已经部分排序（用黑色表示）。
每次迭代时，从输入数据中移除一个元素（用红色表示），并原地将其插入到已排好序的链表中。

插入排序算法：  
1. 插入排序是迭代的，每次只移动一个元素，直到所有元素可以形成一个有序的输出列表。
2. 每次迭代中，插入排序只从输入数据中移除一个待排序的元素，找到它在序列中适当的位置，并将其插入。
3. 重复直到所有输入数据插入完为止。

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

#### 思路
使用一个指针用来标记待排序的元素，一个指针指向待排序的前一个元素。

如果待排序的元素比前一个元素小，则对其进行遍历插入，否则进行下一个比较。

#### 伪代码
```
INSERTION-SORT-LIST(head)
    if head == nil || head.next == nil
        return head
    top = new ListNode()
    top.next = head
    p1, p2 = head, head.next
    for p2 != nil
        if p2.val < p1.val
            temp = p2.next
            p1.next = p2.next
            p3 = top
            for p3.next.val < p2.val
                p3 = p3.next
            p2.next = p3.next
            p3.next = p2
            p2 = temp
        else    
            p1 = p1.next
            p2 = p2.next
    return top.next
```