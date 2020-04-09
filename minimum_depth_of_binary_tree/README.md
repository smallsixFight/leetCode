### 二叉树的最小深度（Minimum Depth of Binary Tree）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/)

#### 题目
给定一个二叉树，找出其最小深度。

最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

说明: 叶子节点是指没有子节点的节点。

示例：
给定二叉树`[3,9,20,null,null,15,7]`，
```
    3
   / \
  9  20
    /  \
   15   7
```
返回它的最小深度 2。

#### 思路
这道题可以用 DFS 或者 BFS 来处理。

#### 题外话
`MinDepthTwo()`为使用 BFS 来处理。

#### 伪代码
```
MIN-DEPTH(root)
    if root == nil
        return 0
    return DFS(root, 0, 1)

DFS(node, high, minHigh)
    high ++
    if node.left == nil && node.right == nil
        if minHigh == 1 || (high > 1 && high < minHigh)
            minHigh = high
        return minHigh
    if node.left != nil
        minHigh = DFS(node.left, high, minHigh)
    if node.right != nil
        minHigh = DFS(node.right, high, minHigh)
    return minHigh
```