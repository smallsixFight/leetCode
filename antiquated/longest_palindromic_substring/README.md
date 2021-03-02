### 最长回文子串（Longest Palindromic Substring）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/longest-palindromic-substring/)

#### 题目
给定一个字符串`s`，找到`s`中最长的回文子串。你可以假设`s`的最大长度为 1000。

示例 1：
```
输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
```
示例 2：
```
输入: "cbbd"
输出: "bb"
```

#### 思路
通过观察可以发现回文具有如下定律：
- 单字符一定是回文；
- 回文的中心为单个字符或双个字符；
- 回文`s[i, j]`中，`s[i+1, j-1]`也一定是回文，即回文字符串一定存在子回文字符串，也可以说，通过回文字符串才能生成更长的回文字符串。

可以发现用动态规划来解决，实现步骤：
- 初始化并记录所有单字符回文和双字符回文的索引值；
- 循环遍历生成更长的回文并添加记录；
- 每次遍历比较并记录最长的回文的长度。

#### 涨姿势
`LongestPalindrome2()`是中心扩展算法的实现过程。

另外，官方还提到了 Manacher 算法。

#### 伪代码
```
LONGEST-PALINDROME(s)
    let record be new queue
    for i = 0 to s.length -1
        ENQUEUE(record, [i, i])
        if i < s.length -1 && s[i] == s[i+1]
            ENQUEUE(record, [i, i+1])
    res = s[0: 1]
    for record.hasNext
        v = DEQUEUE(record)
        if v[0] -1 >= 0 && v[1] +1 < s.length
            if  s[v[0]] == s[v[1]]
                ENQUEUE(record, [v[0] -1, v[1] +1])
                if v[1] - v[0] + 2 > res
                    res = s[v[0]-1: v[1]+1]
    return res
```