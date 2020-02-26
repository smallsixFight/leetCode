### 二叉树的最大深度（Maximum Depth of Binary Tree）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/)

#### 题目
给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

**说明：** 叶子节点是指没有子节点的节点。

示例：
给定二叉树`[3,9,20,null,null,15,7]`，返回它的最大深度 3 。
```
    3
   / \
  9  20
    /  \
   15   7
```

#### 思路
这题可以迭代分别获取左右子树的深度，然后比较返回较大值。

`MaxDepth2()`是基于深度优先探索的写法。

#### 伪代码
```
MAX-DEPTH(root)
    if root == nil
        return 0
    if root.left == nil && root.right == nil
        return 1
    leftDepth = MAX-DEPTH(root.left)
    rightDepth = MAX-DEPTH(root.right)
    if leftDepth > rightDepth
        return leftDepth +1
    else
        return rightDepth +1
```