### 二叉树的右视图（Binary Tree Right Side View）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/binary-tree-right-side-view/)

#### 题目
给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

示例：
```
输入: [1,2,3,null,5,null,4]
输出: [1, 3, 4]
解释:

   1            <---
 /   \
2     3         <---
 \     \
  5     4       <---
```

#### 思路
这道题可以用 BFS 来做。

#### 伪代码
```
RIGHT-SIDE-VIEW(root)
    if root == nil
        return nil
    queue = new Queue()
    queue.enqueue(root)
    res = new Array()
    while !queue.isEmpty()
        l = queue.length
        for i = 0 to l -1
            node = queue.dequeue()
            if node.left != nil
                queue.enqueue(node.left)
            if node.right != nil
                queue.enqueue(node.right)
            if i == l -1
                res.add(node.val)
    return res
```