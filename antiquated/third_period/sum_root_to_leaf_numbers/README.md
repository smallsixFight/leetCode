### 求根到叶子节点数字之和（Sum Root to Leaf Numbers）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/sum-root-to-leaf-numbers/)

#### 题目
给定一个二叉树，它的每个结点都存放一个`0-9`的数字，每条从根到叶子节点的路径都代表一个数字。

例如，从根到叶子节点路径`1->2->3`代表数字`123`。

计算从根到叶子节点生成的所有数字之和。

说明: 叶子节点是指没有子节点的节点。

示例 1：
```
输入: [1,2,3]
    1
   / \
  2   3
输出: 25
解释:
从根到叶子节点路径 1->2 代表数字 12.
从根到叶子节点路径 1->3 代表数字 13.
因此，数字总和 = 12 + 13 = 25.
```
示例 2：
```
输入: [4,9,0,5,1]
    4
   / \
  9   0
 / \
5   1
输出: 1026
解释:
从根到叶子节点路径 4->9->5 代表数字 495.
从根到叶子节点路径 4->9->1 代表数字 491.
从根到叶子节点路径 4->0 代表数字 40.
因此，数字总和 = 495 + 491 + 40 = 1026.
```

#### 思路
这道题应该算是前序遍历的扩展吧，还结合前面做过的一道简单题来着。做前序遍历，直到叶子节点是将数字加入和。

#### 伪代码
```
SUM-NUMBERS(root)
    if root == nil
        return 0
    return DFS(root, 0)

DFS(node, preSum)
    sum = preSum * 10 + node.val - '0'
    if node.left == nil && node.right == nil
        return sum
    if node.left != nil && node.right != nil
        temp = sum
        sum = DFS(node.left, temp)
        sum += DFS(node.right, temp)
    elseif node.left != nil
        sum = DFS(node.left, sum)
    else
        sum = DFS(node.right, sum)
    return sum
```