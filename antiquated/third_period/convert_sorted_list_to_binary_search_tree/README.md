### 有序链表转换二叉搜索树 （Convert Sorted List to Binary Search Tree）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree/)

#### 题目
给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

示例：
```
给定的有序链表： [-10, -3, 0, 5, 9],

一个可能的答案是：[0, -3, 9, -10, null, 5], 它可以表示下面这个高度平衡二叉搜索树：

      0
     / \
   -3   9
   /   /
 -10  5
```

#### 思路
先把链表转换为数组，然后使用二分法找到中间索引的值生成节点，然后继续递归。

#### 伪代码
```
SORTED-LIST-TO-BST(head)
    if head == nil
        return nil
    arr = new Array()
    for head != nil
        arr.push(head.val)
        head = head.next
    return DFS(arr)
    
DFS(arr)
    if arr.length == 0
        return nil
    mid = arr.length /2
    node = new TreeNode(arr[mid])
    node.left = DFS(arr[:i])
    node.right = DFS(arr[i+1:])
    return node
```