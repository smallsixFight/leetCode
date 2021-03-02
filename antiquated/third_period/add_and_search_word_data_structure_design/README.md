### 添加与搜索单词 - 数据结构设计（Add and Search Word - Data structure design）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/add-and-search-word-data-structure-design/)

#### 题目
设计一个支持以下两种操作的数据结构：
```
void addWord(word)
bool search(word)
```
search(word) 可以搜索文字或正则表达式字符串，字符串只包含字母`.`或`a-z`。`.`可以表示任何一个字母。

示例：
```
addWord("bad")
addWord("dad")
addWord("mad")
search("pad") -> false
search("bad") -> true
search(".ad") -> true
search("b..") -> true
```

说明：  
你可以假设所有单词都是由小写字母`a-z`组成的。

#### 思路
这道题跟前面做过的`实现 Trie`的题类似，`search`多加了一个`.`需要匹配，那么就需要利用回溯算法来做，其实也是进行 DFS，不过只要有一个符合条件的就可以直接返回 true。

#### 题外话
我在提交的代码中看到没有使用字典树来实现的，而是把相同长度的字符串存储在同一个数组里面，这不失为一个办法，由于官方的测试案例只有 13 个，所以这个方法在空间和时间上的击败率几乎都是 100%，但我想到如果存在只有尾部字母不同的大量长字符串的时候，空间的使用率就大大增加了，所以没有去写下来。

另外，我一开始使用的 map，改用 slice 后时间缩短了一半，确认了一个新的点。

#### 伪代码

