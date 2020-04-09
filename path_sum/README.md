### 路径总和（Path Sum）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/path-sum/)

#### 题目
给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。

说明: 叶子节点是指没有子节点的节点。

示例：  
给定如下二叉树，以及目标和`sum = 22`，
```
              5
             / \
            4   8
           /   / \
          11  13  4
         /  \      \
        7    2      1
```
返回`true`, 因为存在目标和为 22 的根节点到叶子节点的路径`5->4->11->2`。

#### 思路
这道题可以用 BFS 和 DFS 来解决。

#### 题外话
`HasPathNumTwo()`为使用 DFS 的方法。

一开始，考虑不全，把负数啥的全部限制掉了，导致一开始连续错了好几次。

#### 伪代码
```
HAS-PATH-SUM(root, sum)
    if root == nil
        return false
    if root.val == sum && root.left == nil && root.right == nil
        return true
    stack = new Stack()
    stack.push(root.val)
    l = stack.length
    for l > 0
        for i = 0 to l -1
            node = stack[0]
            stack = stack[1:]
            if node.left != nil
                node.left.val += node.val
                if node.left.val == sum && node.left.left == nil && node.left.right == nil
                    return true
                elseif node.left.val < sum
                    stack.push(node.left.val)
            if node.right != nil
                node.right.val += node.val
                if node.right.val == sum && node.right.left == nil && node.right.right == nil
                    return true
                elseif node.right.val < sum
                    stack.push(node.right.val) 
        l = stack.length
    return false
```