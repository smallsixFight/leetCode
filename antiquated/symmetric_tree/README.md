### 对称二叉树（Symmetric Tree）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/symmetric-tree/)

#### 题目
给定一个二叉树，检查它是否是镜像对称的。

例如，二叉树`[1,2,2,3,4,4,3]`是对称的。
```
    1
   / \
  2   2
 / \ / \
3  4 4  3
```
但是下面这个`[1,2,2,null,3,null,3]`则不是镜像对称的：
```
    1
   / \
  2   2
   \   \
   3    3
```
*说明：*  
如果你可以运用递归和迭代两种方法解决这个问题，会很加分。

#### 思路
跟前面的`Some Tree`这题基本相似，比较变成`root`的递归左右子树，并且是左结点跟右节点比较。

#### 伪代码
```
IS-SYMMETRIC(root)
    if root != nil
        return IS-EQUAL(root.left, root.right)
    return true
    
IS-EQUAL(p, q)
    if p == q           // 均为 nil
        return true
    if p.val == q.val
        return IS-EQUAL(p.left, q.right) && IS-EQUAL(p.right, q.left)
    return false
```