### 不同的二叉树搜索树 II（Unique Binary Search Trees II）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/unique-binary-search-trees-ii/)

#### 题目
给定一个整数 n，生成所有由 1 ... n 为节点所组成的二叉搜索树。

示例：
```
输入: 3
输出:
[
  [1,null,3,2],
  [3,2,null,1],
  [3,1,null,null,2],
  [2,1,3],
  [1,null,2,null,3]
]
解释:
以上的输出对应以下 5 种不同结构的二叉搜索树：

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3
```

#### 思路
手动画几个图就可以发现，第`n`个节点拥有的二叉树取决与第`n-1`个的所有二叉树插入 n 生成的所有二叉树。那么就可以用动态规划来做。

生成新的二叉树操作的描述大致为：对`n-1`的所有二叉树从根节点开始，依次替换右孩子直至成为叶子节点，每次替换后对原子节点进行处理，这就生成新的二叉树了。

#### 伪代码
```
GENERATE-TREES(n)
    res = new Array()
    t1 = new TreeNode(1)
    res.push(t1)
    for i = 2 to n
        l = res.length
        for k = 0 to l
            CREATE0-NEW-TREE(res[k], i)

CREATE-NEW-TREE(tree, n)
    t = tree
    pre = nil
    for t != nil
        newChild = new TreeNode()
        if pre != nil   // 逐个替换右孩子
            temp = t
            pre.right = nil
            newTree = new TreeNode()
            copy(tree, newTree)
            copy(t, newChild)
            node = new TreeNode(n)
            node.left = newChild
            p = newTree
            for p.right != nil
                p = p.right
            p.right = node
            res.push(newTree)
            pre.right = temp
        else    // 作为根节点
            copy(t, newChild)
            node = new TreeNode(n)
            node.left = newChild
            res.push(node)
        pre = t
        t = t.right
    pre.right = new TreeNode(n) // 作为叶子节点
```