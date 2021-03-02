### 填充每个节点的下一个右侧节点指针 II（Populating Next Right Pointers in Each Node II）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/)

#### 题目
给定一个二叉树
```
struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
```
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为`NULL`。

初始状态下，所有 next 指针都被设置为`NULL`。

进阶：

- 你只能使用常量级额外空间。
- 使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。

示例：

![117_sample.png](https://i.loli.net/2020/04/18/MBhckj6RGvHOJeT.png)

```
输入：root = [1,2,3,4,5,null,7]
输出：[1,#,2,3,#,4,5,7,#]
解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B 所示。
```

提示：

- 树中的节点数小于`6000`
- `-100 <= node.val <= 100`

#### 思路
这一题相较于上一题，就是树不是完美二叉树，那么就会有以下的调整：

- 如果没有右孩子，那么左孩子的 next 值是最近有孩子节点的叔节点的左孩子或右孩子；
- 右孩子的 next 值是最近有孩子节点的叔节点的左孩子或右孩子；
- 没有左孩子不一定没有右孩子；
- 需要从右子树先 DFS，因为左子树可能需要查找可用的叔节点。

#### 伪代码
```
CONNECT(root)
    if root == nil
        return nil
    DFS(root, nil)
    return root
    
DFS(node, uncleNode)
    if node.left == nil && node.right == nil
        return
    if node.right != nil
        if uncleNode != nil
            if uncleNode.left != nil
                node.right.next = uncleNode.left
            else
                node.right.next = uncleNode.right
        temp = node.right.next
            for temp != nil && temp.left == nil && temp.right == nil
                temp = temp.next
        DFS(node.right, temp)
    if node.left != nil
        if node.right != nil
            node.left.next = node.right
        else
            if uncleNode.left != nil
                node.left.next = uncleNode.left
            else
                node.left.next = uncleNode.right
        temp = node.left.next
        for temp != nil && temp.left == nil && temp.right == nil
            temp = temp.next
        DFS(node.left, temp)
``` 