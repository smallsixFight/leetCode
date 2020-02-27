### 将有序数组转换为二叉搜索树（Convent Sorted Array to Binary Search Tree）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/)

#### 题目
将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

示例：
```
给定有序数组: [-10,-3,0,5,9],

一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：

      0
     / \
   -3   9
   /   /
 -10  5
```

#### 思路
为了满足高度平衡二叉树这个条件，且由于数组是升序的，`nums[mid]`一直作为树的根结点，可以以此进行二叉树的转换。

#### 伪代码
```
SORTED-ARRAY-TO-BST(nums)
    if nums.length == 0
        return nil
    node = TreeNode{ val: nums[nums.length/2] }
    node.left = SORTED-ARRAY-TO-BST(nums[0:nums.length/2 -1])
    node.right = SORTED-ARRAY-TO-BST(nums[nums.length/2 +1]:)
    return node
```