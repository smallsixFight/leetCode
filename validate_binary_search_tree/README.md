### 验证二叉搜索树（Validate Binary Search Tree）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/validate-binary-search-tree/)

#### 题目
给定一个二叉树，判断其是否是一个有效的二叉搜索树。

假设一个二叉搜索树具有如下特征：

- 节点的左子树只包含小于当前节点的数。
- 节点的右子树只包含大于当前节点的数。
- 所有左子树和右子树自身必须也是二叉搜索树。

示例 1：
```
输入:
    2
   / \
  1   3
输出: true
```
示例 2：
```
输入:
    5
   / \
  1   4
     / \
    3   6
输出: false
解释: 输入为: [5,1,4,null,null,3,6]。
     根节点的值为 5 ，但是其右子节点值为 4 。
```

#### 思路
这道题可以使用递归来实现一个自底向上的方法，每次递归返回三个值`isOk`、`minVal`和`maxVal`。如果`isOk == false`或`node.val >= maxVal || node.val <= minVal`，则直接返回 false。

minVal 和 maxVal 分别代表子树的最小值和最大值，isOk 表示子树是否为符合标准的 BST。

#### 思路二
其实也是自顶向下，先讲根结点作为最大值和最小值，然后分别对左右子树进行 DFS，迭代根据左右子树更换最大值和最小值，然后返回结果。

#### 伪代码
```
IS-VALID-BST(root)
    if root == nil
        return true
    isOk, _, _, = DFS(root)
    return isOk
    
DFS(node)
    if node.left != nil && node.left.val >= node.val
        return false
    if node.right != nil && node.right.val <= node.val
        return false
    minV, maxV = node.val, node.val
    if node.left != nil
        isOk, minVal, maxVal = DFS(node.left)
        if !isOk || maxVal >= node.val
            return false, 0, 0
        minV = minVal
    if node.right != nil
        isOk, minVal, maxVal = DFS(node.right)
        if !isOk || minVal <= node.val
            return false
        maxV = maxVal
    return true, minVal, maxVal
```