### 串联所有单词的子串（Substring with Concatenation of All Words）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words/)

#### 题目
给定一个字符串`s`和一些长度相同的单词`words`。找出`s`中恰好可以由`words`中所有单词串联形成的子串的起始位置。

注意子串要与`words`中的单词完全匹配，中间不能有其他字符，但不需要考虑`words`中单词串联的顺序。

示例 1：
```
输入：
  s = "barfoothefoobarman",
  words = ["foo","bar"]
输出：[0,9]
解释：
从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
输出的顺序不重要, [9,0] 也是有效答案。
```
示例 2：
```
输入：
  s = "wordgoodgoodgoodbestword",
  words = ["word","good","best","word"]
输出：[]
```

#### 思路一
这道题让我想起《算法导论》里面贪心算法的第一个例子。

先把`s`串中能匹配到的单词起始索引存起来，并且记录每个单词重复的个数。

然后从头开始遍历，如果能够生成一个符合要求的等差数列，记录等差数列的起始位置。

#### 思路二
思路一提交后运行的时间超过 800ms，简直无颜，然后在官方题解看到滑动窗口的方法，然后做了整理和做了实现，大致思路如下：

- 前面部分跟思路一类似，先使用一个 hash table 记录单词数组的单词和出现的次数；
- 当找到符合条件的首个单词时，记录起始位置，然后使用一个计数器来记录当前符合条件的单词总数；
- 此外，由于单词的重复个数不能超过`words`中的个数，所以需要一个数组`remain`来记录当前符合要求的单词及次数，当超过限制的次数则不符合条件；
- 当存在符合的单词而单词次数过载时，滑动窗口向后移动一个单词的长度，如果还是存在前面的情况，继续向后移动窗口，直到不再过载，那么继续向后匹配；
- 当计数器的数跟`words`长度一样时，添加到结果集并让窗口向后移动；
- 当单词不不匹配时，如果计数器`count`的值不为 0，那么需要重置`remain`和`count`。

#### 题外话
伪代码的`FIND-SUBSTRING-TWO`实际上是优化后的代码，实际实现代码为`FindSubstringThree()`。

思路二的最后一条如果没有做到，即使是优化后的代码`FindSubstringThree()`时间也会花费 64ms，没优化的代码`FindSubstringTwo()`会花费长达 300ms 多。

一开始很头痛，一直以为滑窗方法没写对，后台通过`pprof`进行分析，才发现有很多不必要的重置操作。

#### 伪代码
```
FIND-SUBSTRING-TWO(s, words)
    res = new Integer Array()
    if s.length == 0 || words.length == 0
        return res
    dupMap = new Map()
    for i = 0 to words.length -1
        dupMap[words[i]] ++
    wordLen = words[0].length
    for i = 0 to wordLen -1
        left, right = i, i
        count = 0
        remainMap = new Map()
        while right + wordLen <= s.length
            if _, ok := dupMap[s[right: right+wordLen]]; ok
                remainMap[s[right: right+wordLen]] ++
                count ++
                while remainMap[s[right: right+wordLen]] > dupMap[s[right: right+wordLen]]
                    remainMap[left: left+wordLen] --
                    left += wordLen
                    count --
            else
                left += wordLen
                right = left
                remainMap = new Map()
                count = 0
                continue
            right += wordLen
            if count == words.length
                res.push(left)
    return res

FIND-SUBSTRING(s, words)
    res = new Integer Array()
    if s.length == 0 || words.length == 0
        return res
    dupMap = new Map()
    for i = 0 to words.length -1
        dupMap[words[i]] ++
    wordLen = words[0].length
    recordMap = new Map()
    for i = 0 to s.length -wordLen
        for j = 0 to words.length -1
            if s[i: i+wordLen] == words[j]
                recordMap[i] = words[j]
                break
    for k := range recordMap
        count = 0
        tempMap = new Map()
        i := 0
        while k + i < s.length && count < words.length
            val, ok = recordMap[k+i]
            if ok
                tempMap[val]  ++
                if tempMap[val] > dupMap[val]
                    break
                count ++
                continue
            break
        if count == words.length
            res.push(k)
    return res
```