### 验证回文串（Valid Palindrome）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/valid-palindrome/)

#### 题目
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

说明：本题中，我们将空字符串定义为有效的回文串。

示例 1：
```
输入: "A man, a plan, a canal: Panama"
输出: true
```
示例 2：
```
输入: "race a car"
输出: false
```

#### 思路
这道题可以用双指针法来做，需要忽略除字母和数字以外的字符。

#### 伪代码
IS-PALINDROME(s)
    if s.length < 2
        return true
    p1, p2 = 0, s.length -1
    while true
        while s[p1] <  48 || (s[p1] > 57 && s[p1] < 65) || (s[p1] > 90 && s[p1] < 97) || s[p1] > 122
            p1 ++
            if p1 == s.length
                break
        while s[p2] <  48 || (s[p2] > 57 && s[p2] < 65) || (s[p2] > 90 && s[p2] < 97) || s[p2] > 122
            p2 --
            if p2 == 0
                break
        if p1 >= p2
            break
        if s[p1] > 96
            s[p1] = s[p1] - 32
        if s[p2] > 96
            s[p2] = s[p2] - 32
        if s[p1] != s[p2]
            return false
    return true