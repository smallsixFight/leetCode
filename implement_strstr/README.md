### 实现`strStr()`（Implement strStr()）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/implement-strstr/)

#### 题目
给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

示例 1：
```
输入: haystack = "hello", needle = "ll"
输出: 2
```
示例 2：
```
输入: haystack = "aaaaa", needle = "bba"
输出: -1
```

说明：  
当`needle`是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。

对于本题而言，当`needle`是空字符串时我们应当返回 0 。这与 C 语言的 strstr() 以及 Java 的 indexOf() 定义相符。

#### 思路
这题可以使用 KMP 算法进行匹配比较，可以降低指针回溯的位置，降低比较的时间。  
首先明确下临界：  
1. haystack 的长度必须不小于 needle 的长度；
2. needle 为空字符串时直接返回 0。

KMP 算法和核心是：  
1. 主串不需要回溯；
2. 子串尽可能回溯到有效的匹配位置。

所以 KMP 需要一个额外的空间为 O(needle.length)来保存子串回溯的位置。

#### 学习到思路
Strstr2() 是我在 LeetCode 看到的另一份解，用的是 Sunday 算法，这个答案充分利用了 golang 中 slice 的特性，自愧不如。

#### 伪代码
```
IMPLEMENT-STRSTR(haystack, needle)
    if haystack.length < needle.length
        return -1
    if needle.length == 0
        return 0
    next = GET-NEXT(needle)
    i, j = 0, 0
    while i < haystack.length && j < needle.length
        if j == -1 || haystack[i] == needle[j]
            i ++
            j ++
        else
            j = next[j]
    if j == needle.length -1
        return i - j
    return -1

GET-NEXT(s)
     next = new array()
     next[0] = -1
     k, j = -1, 0
     for j to s.length -1
        if k == -1 || s[k] == s[j]
            if s[++k] == s[++j]
                next[j] = next[k]
            else
                next[j] = k
        else
            k = next[k]
    return next
 ```