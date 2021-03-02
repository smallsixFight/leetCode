### 二叉树的层次遍历 II（Binary Tree Level Order Traversal II）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/)

#### 题目
给定一个二叉树，返回其节点值自底向上的层次遍历（即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）。

示例：
给定二叉树`[3,9,20,null,null,15,7]`，
```
    3
   / \
  9  20
    /  \
   15   7
```
返回其自底向上的层次遍历为：
```
[
  [15,7],
  [9,20],
  [3]
]
```

#### 思路
这题可以使用深度优先探索来完成；把不同深度的值存储在一起，再逆序生成返回值。

#### 伪代码
```
LEVEL-ORDER-BOTTOM(root)
    record = new [][]int
    DFS-VISIT(root, 0, record)
    i, j = 0, record.length -1
    for i < j
        temp = record[i]
        record[i] = record[j]
        record[j] = temp
        i ++
        j --
    return record

DFS-VISIT(node, depth, record)
    if node != nil
        record[depth].add(node.val)
        DFS-VISIT(node.left, depth +1, record)
        DFS-VISIT(node.right, depth +1, record)
```     