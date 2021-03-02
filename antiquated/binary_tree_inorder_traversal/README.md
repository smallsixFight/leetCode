### 二叉树的中序遍历（Binary Tree Inorder Traversal）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/)

#### 题目
给定一个二叉树，返回它的`中序`遍历。

示例：
```
输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,3,2]
```

**进阶：** 递归算法很简单，你可以通过迭代算法完成吗？

#### 思路
这题我直接写下迭代的思路。递归其实是函数的入栈出栈，使得可以回到原来的节点进行后续操作。

那么迭代的话，就是要自己把栈给提炼出来。大致思路如下：

- 从根节点开始，如果有左孩子，则入栈，指向左孩子；
- 如果没有左孩子，将 val 保存到结果集，然后指向右孩子；
- 如果没有右孩子，则取出栈顶的一个元素，从栈取出的元素不用再去判断左孩子。

#### 伪代码
```
INORDER-TRAVERSAL(root)
    res = new Array()
    if root == nil
        return res
    stack = new Stack()
    p = root
    while stack.length > 0 || p != nil
        isPop = false
        if p == nil
            p = stack.pop()
            isPop = true
        if !isPop && p.Left != nil
            stack.push(p)
            p = p.left
            continue
        res.push(p.val)
        if p.right != nil
            p = p.right
        else
            p = nil
    return res
```