### 正则表达式匹配（Regular Expression Matching）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/regular-expression-matching/)

#### 题目
给你一个字符串`s`和一个字符规律`p`，请你来实现一个支持`'.'`和`'*'`的正则表达式匹配。
```
'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
```
所谓匹配，是要涵盖**整个**字符串`s`的，而不是部分字符串。

**说明：**  
- `s`可能为空，且只包含从`a-z`的小写字母。
- `p`可能为空，且只包含从`a-z`的小写字母，以及字符`.`和`*`。

示例 1：
```
输入:
s = "aa"
p = "a"
输出: false
解释: "a" 无法匹配 "aa" 整个字符串。
```
示例 2：
```
输入:
s = "aa"
p = "a*"
输出: true
解释: 因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
```
示例 3：
```
输入:
s = "ab"
p = ".*"
输出: true
解释: ".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
```
示例 4：
```
输入:
s = "aab"
p = "c*a*b"
输出: true
解释: 因为 '*' 表示零个或多个，这里 'c' 为 0 个, 'a' 被重复一次。因此可以匹配字符串 "aab"。
```
示例 5：
```
输入:
s = "mississippi"
p = "mis*is*p*."
输出: false
```

#### 思路一
这道题主要是对于`*`的处理，`*`可以匹配前面零个或多个相同的字符，即可以得到下面的条件：

- `*`后面的模式串跟匹配串能够匹配成功即可达到目的（零个匹配的情况下）；
- 匹配串除去当前匹配成功的字符，剩余的匹配串能跟模式串匹配成功即可达到目的（一个或多个匹配的情况下）；

以上两种有一种成功即意味着匹配串和模式串能够成功匹配，那么可以使用回溯法来做。

#### 思路二
第二个思路是看到官方题解有用到动态规划，于是对思路一进行整理，可以发现：当存在`p[1]`为`*`的时候，存在求重复子问题的解的情况，那么可以将子问题的解记录起来减少求解的次数。

重复子问题是——求`s`和`p`的前`i`个和前`j`个是否匹配，所以可以使用一个二维的数据来记录。

#### 题外话
这道题主要是能够认知到对`*`的处理方式，我一开始只按照单个方向去匹配，写了很长又复杂的代码，还是没有达到条件。

#### 伪代码
```
// 回溯法
IS-MATCH(s, p)
    if p.length == 0
        return s.length == 0
    firstMatch = s.length != 0 && (s[0] == p[0] || p[0] == '.')
    if p.length > 1 && p[1] == '*'
        return IS-MATCH(s, p[2:]) || (firstMatch && IS-MATCH(s[1:], p))
    else
        return firstMatch && IS-MATCH(s[1:], p[1:])
        
// DP
IS-MATCH(s, p)
    if p.length == 0
        return s.length == 0
    dp = new array
    dp[s.length][p.length] = true
    for i = s.length to 0
        for j = p.length -1 to 0
            firstMatch = i < s.length && (s[i] == p[j] || p[j] == '.')
            if j +1 < p.length && p[j+1] == '*'
                dp[i][j] = dp[i][j+2] || (firstMatch && dp[i+1][j])
            else
                dp[i][j] == firstMatch && dp[i+1][j+1]
    return dp[0][0]
```