### 实现 Trie (前缀树)（Implement Trie (Prefix Tree)）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/implement-trie-prefix-tree/)

#### 题目
实现一个 Trie (前缀树)，包含`insert`，`search`和`startsWith`这三个操作。

示例：
```
Trie trie = new Trie();

trie.insert("apple");
trie.search("apple");   // 返回 true
trie.search("app");     // 返回 false
trie.startsWith("app"); // 返回 true
trie.insert("app");   
trie.search("app");     // 返回 true
```

说明：
- 你可以假设所有的输入都是由小写字母`a-z`构成的。
- 保证所有输入均为非空字符串。

#### 思路
思路大致就是把字符串的每个字符单独存储起来。`insert()`当遇到当前位置不存在的字符时，则扩展一个并加入到父节点中。

为了方便判断，我使用 map 来存储，另外使用一个 map 来保存新增的字符串，为键查找时匹配返回结果。

这道题依旧是涉及设计相关的题目，还是直接看代码，不写伪代码了。

#### 题外话
根据官方的题解，我把保存字符串的 map 去掉，添加了 isEnd 的标志，所以在键查找需要进行遍历，节省了近 4M 的空间。

#### 伪代码
