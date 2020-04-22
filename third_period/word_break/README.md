### 单词拆分（Word Break）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/word-break/)

#### 题目
给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。

说明：

- 拆分时可以重复使用字典中的单词。
- 你可以假设字典中没有重复的单词。

示例 1：
```
输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
```
示例 2：
```
输入: s = "applepenapple", wordDict = ["apple", "pen"]
输出: true
解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
     注意你可以重复使用字典中的单词。
```
示例 3：
```
输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
输出: false
```

#### 思路
先将字典转换为一个 hashMap，然后用一个数组来记录能够拆分的末位置，然后遍历字符串，每个位置的字符和数组的任意一个之间的单词能够组成字典中的单词，那么就可以可以拆分的。

#### 伪代码
```
WORD-BREAK(s, wordDict)
    if s == "" || wordDict.length == 0
        return false
    m = new Map()
    for i = 0 to wordDict.length -1
        m[wordDict[i]] = true
    if s.length == 1
        return m[s]
    m2 = new Map()
    m2[-1] = true   // 用于匹配整个单词是否存在
    for i = 0 to s.length -1
        for k in m2
            if m[s[k+1:i+1]]
                m2[i] = true
                break
    return m2[s.length-1]
```