### 二叉树的锯齿形层次遍历（Binary Tree Zigzag Level Order Traversal）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/)

#### 题目
给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

示例：  
给定二叉树`[3,9,20,null,null,15,7]`，
```
    3
   / \
  9  20
    /  \
   15   7
```
返回锯齿形层次遍历如下：
```
[
  [3],
  [20,9],
  [15,7]
]
```

#### 思路
这道题也不难，只要基于`二叉树的层次遍历`的 BFS 方法稍微修改一下，加一个是否逆序的标记就行。

#### 伪代码
```
ZIGZAG-LEVEL-ORDER(root)
    if root == nil
        return nil
    res = new Array()
    record = new Array()
    record.push(root)
    isReverse = false
    for record.length > 0
        arr = new Array()
        l = record.length
        for i = l -1 to 0
            arr.push(record[i].val)
            if isReverse
                if record[i].right != nil
                    record.push(record[i].right)
                if record[i].left != nil
                    record.push(record[i].left)
            else
                if record[i].left != nil
                    record.push(record[i].left)
                if record[i].right != nil
                    record.push(record[i].right)
        res.push(arr)
        record = record[l:]
        isReverse = !isReverse
    return res
```