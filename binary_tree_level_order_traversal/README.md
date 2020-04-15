### 二叉树的层次遍历（Binary Tree Level Order Traversal）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/binary-tree-level-order-traversal/)

#### 题目
给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）

示例：  
二叉树：`[3,9,20,null,null,15,7]`，
```
    3
   / \
  9  20
    /  \
   15   7
```
返回其层次遍历结果：
```
[
  [3],
  [9,20],
  [15,7]
]
```

#### 思路
先前做过几乎相同的题目，而且只是变成自顶向下层级遍历而已，代码可以直接拉过来用。

另外再写一个 BFS 的实现。

#### 伪代码
```
LEVEL-ORDER(root)
    record = new [][]int
    DFS-VISIT(root, 0, record)
    i, j = 0, record.length -1
    return record

DFS-VISIT(node, depth, record)
    if node != nil
        record[depth].add(node.val)
        DFS-VISIT(node.left, depth +1, record)
        DFS-VISIT(node.right, depth +1, record)
```     
