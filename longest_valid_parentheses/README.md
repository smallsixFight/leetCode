### 最长有效括号（Longest Valid Parentheses）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/longest-valid-parentheses/)

#### 题目
给定一个只包含`'('`和`')'`的字符串，找出最长的包含有效括号的子串的长度。

示例 1：
```
输入: "(()"
输出: 2
解释: 最长有效括号子串为 "()"
```
示例 2：
```
输入: ")()())"
输出: 4
解释: 最长有效括号子串为 "()()"
```

#### 思路
这道题求的是最长有效括号的长度，而不是有效括号的个数（一开始看错还吐槽这居然是困难题）。有如下一些条件：

- 只有当前为`)`才需要尝试向前匹配`(`，向前匹配需要减去`n-1`的有效括号的长度；
- 如果能够匹配得到，那么当前的有效长度 +2；另外，匹配到的不是`s[0]`，需要跟前面的有效长度连接起来。

#### 伪代码
```
LONGEST-VALID-PARENTHESES(s)
    if s.length < 2
        return 0
    dp = new Array()
    maxLen = 0
    for i = 0 to s.length -1
        dp.push(0)
    for i =1 to s.length -1
        if s[i] == ')'
            idx = i -dp[i-1] -1
            if idx >= 0 && s[idx] == '('
                dp[i] = dp[i-1] +2
                if idx > 0 // 连接
                    dp[i] += dp[idx-1]
                if dp[i] > maxLen
                    maxLen = dp[i]
    return maxLen
```