### 最长公共前缀（Longest Common Prefix）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/longest-common-prefix/)

#### 题目
编写一个函数来查找字符串数组中的最长公共前缀。  
如果不存在公共前缀，返回空字符串 ""。

示例1：
```
输入: ["flower","flow","flight"]
输出: "fl"
```
示例2：
```
输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
```

#### 思路
这题可以用贪心思想来解决。因为求的是公共前缀，所以字符串数组`strs`中的 s1 和 s2 的最长公共前缀最优解一定属于`strs`的最优解。  
另外，如果前 n 个的最优解是空字符串，则可以直接返回。

#### 伪代码
```
LONGEST-COMMON-PREFIX(strs)
    if strs.length < 2
        return ""
    result = strs[0]
    for i = 1 to strs.length -1
        result = COMMON-PREFIX(result, strs[i])
        if result == ""
            break
    return result
    
COMMON-PREFIX(s1, s2)
    min = math.Min(s1.length, s2.length)
    rv = ""
    for i = 0 to min -1
        if s1[i] == s2[i]
            rv += s1[i]
        else return rv
    return rv
```