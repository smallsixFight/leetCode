### 二叉树的前序遍历（Binary Tree Preorder Traversal）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/binary-tree-preorder-traversal/)

#### 题目
给定一个二叉树，返回它的 前序 遍历。

示例：
```
输入: [1,null,2,3]  
   1
    \
     2
    /
   3 

输出: [1,2,3]
```

进阶：递归算法很简单，你可以通过迭代算法完成吗？

#### 思路
这道题可以用 DFS 来做。前面好像做过中序的迭代实现。

#### 伪代码
```
PREORDER-TRAVERSAL(root)
    if root == nil
        return root
    res = new Array()
    stack = new Stack()
    stack.push(root)
    while !stack.isEmpty()
        node = stack.pop()
        res.push(node.val)
        if node.right != nil
            stack.push(node.right)
        if node.left != nil
            stack.push(node.left)
    return res
```