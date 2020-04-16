### 路径总和 II（Path Sum II）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/path-sum-ii/)

#### 题目
给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。

说明: 叶子节点是指没有子节点的节点。

示例：  
给定如下二叉树，以及目标和`sum = 22`，
```
              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
```
返回：
```
[
   [5,4,11,2],
   [5,8,4,5]
]
```

#### 思路
之前做过 I 吧，不过已经忘记了，这道题的要求是要从根结点到叶子节点，所以靠遍历到叶子节点才能判断是否能够加入结果集。

这道题用 DFS 来做即可。

#### 题外话
第一次提交代码错误后，我马上发现我又犯了跟做 I 那道题一样的错误，没去考虑负数。

#### 伪代码
```
PATH-SUM(root, sum)
    res = new array()
    return DFS(node, sum, new Array(), res)

DFS(node, sum, arr, res)
    if node == nil || (node.val != sum && node.left == nil && node.right == nil)
        return
    arr.push(node.val)
    if node.val == sum && (node.left == nil && node.right == nil)
        res.push(arr)
    DFS(node.left, sum - node.val, arr, res)
    copy(arr, newArr)
    DFS(node.right, sum - node.val, newArr, res)
```