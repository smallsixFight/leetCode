### 分割回文串（Palindrome Partitioning）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/palindrome-partitioning/)

#### 题目
给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。

返回 s 所有可能的分割方案。

示例：
```
输入: "aab"
输出:
[
  ["aa","b"],
  ["a","a","b"]
]
```

#### 思路
这道题拿过来就能看出需要用回溯法来做。

先找出第一个回文字符串，然后在剩下的字符串找下一个回文字符串，剩下的字符串能够刚好组成一个结果集，则保存，否则则回溯，进行下一个回文串查找和判断。

为了避免重复的回文判断，可以先将所有符合条件的回文串的索引保存起来。

#### 伪代码
```
PARTITION(s)
    record = new Map()
    for i = 0 to s.length -1
        for j = i to s.length -1
            IS-PALINDROME(i, j, record)
    res = new Array()
    for i = 0 to str.length -1
        DFS(record, new Array(), 0, i)

DFS(record, arr, a, b)
    if a == str.length || b == str.length  || !record[]int{a,b
        return
    arr.push(str[a:b+1])
    if b == str.length -1
        temp = arr
        res.push(temp)
        return
    for i = b+1 to str.length -1
        DFS(record, arr, b+1, i)

IS-PALINDROME(startIdx, endIdx, record)
    sa, ea = startIdx, endIdx
    while startIdx <= endIdx
        if s[startIdx] != s[endIdx]
            return
        startIdx ++
        endIdx --
    record.push([]int{sa, ea)
```