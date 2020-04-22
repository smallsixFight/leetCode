### 环形链表 II（Linked List Cycle II）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/linked-list-cycle-ii/)

#### 题目
给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回`null`。

为了表示给定链表中的环，我们使用整数`pos`来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果`pos`是`-1`，则在该链表中没有环。

说明：不允许修改给定的链表。

示例 1：
```
输入：head = [3,2,0,-4], pos = 1
输出：tail connects to node index 1
解释：链表中有一个环，其尾部连接到第二个节点。
```
![circularlinkedlist.png](https://i.loli.net/2020/04/22/Z8CIKDrwVGihBHR.png)
示例 2：
```
输入：head = [1,2], pos = 0
输出：tail connects to node index 0
解释：链表中有一个环，其尾部连接到第一个节点。
```
![circularlinkedlist_test2.png](https://i.loli.net/2020/04/22/i129lesPDtnRc4F.png)
示例 3：
```
输入：head = [1], pos = -1
输出：no cycle
解释：链表中没有环。
```
![circularlinkedlist_test3.png](https://i.loli.net/2020/04/22/DnIAzXW2RcdHaBb.png)

进阶：  
你是否可以不用额外空间解决此题？

#### 思路一
使用一个 hashMap 轻松解决。

#### 思路二
不知道怎么想到中学时常做的相遇问题，然后以此扩展出来的思路。首先，要确保存在环，可以使用两个指针，一个每次循环向后移动一个单位，一个向后移动两个单位，如果能够相遇，则存在环，且环在相遇的位置或前面的位置。

再次使用两个指针开始遍历，一个指针从头向后移动，每次移动第二个指针都要绕一圈，如果第二指针回到相遇的位置，则第一个指针的位置不是入口，向后移动以为，继续遍历；

如果第一个指针移动到相遇的位置，则相遇的位置是入口点。

进一步优化，相遇点让第一阶段的两个指针距离入口点的距离相同，可以得出，入口点到 head 的距离等于第二个指针从相遇点绕圈回到入口点的距离。

#### 伪代码
```
DETECT-CYCLE(head)
    if head == nil
        return head
    p1, p2 = head, head.next
    for p1 != p2
        if p2 != nil && p2.next != nil
            p2 = p2.next.next
        else
            return nil
        p1 = p1.next
    p1 = head
    p3 = p2
    for p1 != p2
        if p3 == p1
            return p1
        p3 = p3.next
        if p3 == p2
            p1 = p1.next
    return p1
```