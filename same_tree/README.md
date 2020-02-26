### 相同的树（Same Tree）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/same-tree/)

#### 题目
给定两个二叉树，编写一个函数来检验它们是否相同。

如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。

示例 1：
```
输入:       1         1
          / \       / \
         2   3     2   3

        [1,2,3],   [1,2,3]

输出: true
```
示例 2：
```
输入:      1          1
          /           \
         2             2

        [1,2],     [1,null,2]

输出: false
```
示例 3：
```
输入:       1         1
          / \       / \
         2   1     1   2

        [1,2,1],   [1,1,2]

输出: false
```

#### 思路
这题可以使用前序查找比较来实现。

#### 伪代码
```
IS-SAME-TREE(p, q)
    if p == nil && q == nil
        return true
    if (p == nil && q != nil) || (p != nil && q == nil) || p.val != q.val
        return false
    return IS-SAME-TREE(p.left, q.left) && IS-SAME-TREE(p.right, q.right)
```