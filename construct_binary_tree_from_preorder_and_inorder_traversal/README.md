### 从前序与中序遍历序列构造二叉树（Construct Binary Tree from Preorder and Inorder Traversal）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/)

#### 题目
根据一棵树的前序遍历与中序遍历构造二叉树。

注意：  
你可以假设树中没有重复的元素。

例如，给出
```
前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
```
返回如下的二叉树：
```
    3
   / \
  9  20
    /  \
   15   7
```

#### 思路
其实可以发现一些规律：

- 前序排列的第一个元素必定为根元素；
- 中序排序根元素前面的元素为左子树元素，后面为右子树元素；
- 中序根元素前面的元素数量和数字跟前序排序根元素的后面相同数量的元素相同；后面的与前序排序剩下的元素相同。

#### 题外话
最近几道二叉树的题目让我感觉回到大学上《数据结构》的时候（0^0）。

#### 伪代码
```
BUILD-TREE(preorder, inorder)
    if preorder.length == 0
        return nil
    idx = 0
    node = new TreeNode(preorder[0])
    for i = 0 to inorder.length -1
        if inorder[i] == preorder[0]
            idx = i
            break
    node.left = BUILD-TREE(preorder[1:idx+1], inorder[:idx])
    node.right = BUILD-TREE(preorder[idx+1:], inorder[idx+1:])
    return node
```