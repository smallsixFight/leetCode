### 从中序与后序遍历序列构造二叉树（Construct Binary Tree from Inorder and Postorder Traversal）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/)

#### 题目
根据一棵树的中序遍历与后序遍历构造二叉树。

注意：  
你可以假设树中没有重复的元素。

例如，给出
```
中序遍历 inorder = [9,3,15,20,7]
后序遍历 postorder = [9,15,7,20,3]
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
这道题跟前面的那道题类似，只是前序变为后序，即`[左子树][右子树][根结点]`，把前面的代码稍微变化一下就好了。

#### 伪代码
```
BUILD-TREE(inorder, postorder)
    if postorder.length == 0
        return nil
    idx = 0
    l = postorder.length
    node = new TreeNode(postorder[l-1])
    for i = 0 to inorder.length -1
        if inorder[i] == postorder[l-1]
            idx = i
            break
    node.left = BUILD-ORDER(inorder[:idx], postorder[:idx])
    node.roght = BUILD-ORDER(inorder[idx+1:], postorder[idx:l-1])
    return node
```