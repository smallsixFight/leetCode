### 二叉树展开为链表（Flatten Binary Tree to Linked List）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/)

#### 题目
给定一个二叉树，**原地**将它展开为链表。

例如，给定二叉树
```
    1
   / \
  2   5
 / \   \
3   4   6
```
将其展开为：
```
1
 \
  2
   \
    3
     \
      4
       \
        5
         \
          6
```

#### 思路
这道题其实就是前序遍历的扩展。


#### 伪代码
```
FLATTEN(root)
    if root == nil
        return nil
    DFS(root)
    
DFS(node)
    p1, rp = nil
    if node.left != nil && node.right != nil
        rp = node.right
    if node.left != nil
        node.right = node.left
        p1 = DFS(node.left)
    else if node.right != nil
        DFS(node.right)
    if rp != nil
        if p1 != nil
            for p1.right != nil
                p1 = p1.right
            p1.right = rp
            DFS(rp)
        else
            node.left.right = rp
    node.left = nil
    return node
```