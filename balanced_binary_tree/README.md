### 平衡二叉树（Balanced Binary Tree）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/balanced-binary-tree/)

#### 题目
给定一个二叉树，判断它是否是高度平衡的二叉树。

本题中，一棵高度平衡二叉树定义为：

> 一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。

示例 1：  
给定二叉树`[3,9,20,null,null,15,7]`
```
    3
   / \
  9  20
    /  \
   15   7
```
返回`true`。

示例 2：  
给定二叉树`[1,2,2,3,3,null,null,4,4]`
```
       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
```
返回`false`。

#### 思路
对根结点的左右子树分别使用 DFS，只有存在一个节点的左右子树高度差大于 1，则整棵树都不是高度平衡树。

只要一个节点的左孩子或右孩子为 nil，那么判断一下他的左子树或者右子树高度是否大于 1 即可。

额，我这个是自底向上的方式，只有子树是平衡的采取判断当前树是否平衡。

#### 伪代码
```
IS-BALANCED(root)
    if root == nil
        return true
    return DFS(root)

DFS(node)
    if node == nil
        return true, 0
    lh, rh = 1, 1
    isLB, h1 = DFS(node.left)
    if !isLB
        return false, 0
    lh += h1
    isRB, h2 = DFS(node.right)
    if !isRB
        return false, 0
    rh += h2
    if math.Abs(lh-rh) > 1
        return false, 0
    return true, max(lh, rh)
```