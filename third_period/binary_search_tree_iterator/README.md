### 二叉搜索树迭代器（Binary Search Tree Iterator）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/binary-search-tree-iterator/)

#### 题目
实现一个二叉搜索树迭代器。你将使用二叉搜索树的根节点初始化迭代器。

调用`next()`将返回二叉搜索树中的下一个最小的数。

示例：

![bst-tree.png](https://i.loli.net/2020/04/29/5WsqvPLGIB9gMJ6.png)
```
BSTIterator iterator = new BSTIterator(root);
iterator.next();    // 返回 3
iterator.next();    // 返回 7
iterator.hasNext(); // 返回 true
iterator.next();    // 返回 9
iterator.hasNext(); // 返回 true
iterator.next();    // 返回 15
iterator.hasNext(); // 返回 true
iterator.next();    // 返回 20
iterator.hasNext(); // 返回 false
```

提示：

- `next()`和`hasNext()`操作的时间复杂度是 O(1)，并使用 O(h) 内存，其中 h 是树的高度。
- 你可以假设`next()`调用总是有效的，也就是说，当调用`next()`时，BST 中至少存在一个下一个最小的数。

#### 思路
这道题在于对栈、中序遍历、 BST 的性质的了解和应用。先将 root 的左孩子入栈，当调用 next() 时出栈，然后判断出栈的元素是否有右孩子，有的话，入栈，并循环将它的左孩子入栈。

#### 题外话
这道题其实涉及到我个人不太熟悉的`分摊复杂度`，我看《算法导论》没看这一章，虽然`next()`函数有个循环遍历，但是 BST 的节点数量是一定的，且当某个节点进行 n 次遍历，那么就会有 n 个节点不需要遍历，所以均摊后复杂度依然是 O(1)。这里纯粹个人理解。

#### 伪代码
代码很简单，主要是构造函数和 next() 函数的实现。
