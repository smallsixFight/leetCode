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

#### 伪代码

