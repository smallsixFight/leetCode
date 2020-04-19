### 单词接龙（Word Ladder）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/word-ladder/)

#### 题目
给定两个单词（beginWord 和 endWord）和一个字典，找到从 beginWord 到 endWord 的最短转换序列的长度。转换需遵循如下规则：
1. 每次转换只能改变一个字母。
2. 转换过程中的中间单词必须是字典中的单词。

说明：

- 如果不存在这样的转换序列，返回 0。
- 所有单词具有相同的长度。
- 所有单词只由小写字母组成。
- 字典中不存在重复的单词。
- 你可以假设 beginWord 和 endWord 是非空的，且二者不相同。

示例 1：
```
输入:
beginWord = "hit",
endWord = "cog",
wordList = ["hot","dot","dog","lot","log","cog"]

输出: 5

解释: 一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog",
     返回它的长度 5。
```
示例 2：
```
输入:
beginWord = "hit"
endWord = "cog"
wordList = ["hot","dot","dog","lot","log"]

输出: 0

解释: endWord "cog" 不在字典中，所以无法进行转换。
```

#### 思路
这一题的目标在于需要基于 BFS 将所有可以尝试的转换都尝试一遍，如果成功转换，那么返回序列长度。

#### 题外话
一开始我没有使用通用状态记录，然后提交后超时了，没想到想节约空间反倒超时了，没能把握题目要求的平衡度（0.0）。

#### 伪代码
```
LADDER-LENGTH(beginWord, endWord, wordList)
    isExist = false
    // 将字典转换为 hash table，用来判断是否已经使用过
    m = new Map()
    commonStatus = new Map()
    for i = 0 to wordList.length -1
        if wordList[i] == endWord
            isExist = true
        if wordList[i] == beginWord
            continue
        for j = 0 to wordList[i].length -1
            newWord = wordList[i][:j] + "*" + wordList[i][j+1:]
            commonStatus[newWord].push(wordList[i])
        m[wordList[i]] = false
    if !isExist
        return 0
    queue = new Array()
    queue.push(beginWord)
    res = -1
    for queue.length > 0
        res += 1
        l = queue.length
        for i = 0 to l -1
            if queue[i] == endWord
                return res +1
            for j = 0 to beginWord.length -1
                newWord = wordList[i][:j] + "*" + wordList[i][j+1:]
                if commonStatus.hasKey(newWord)
                    v = commonStatus.get(newWord)
                    for k = 0 to v.length -1
                        if !m[v[k]]
                            queue.push(v[k])
                            m[v[k]] = true
        queue = queue[l:]
    return 0
```