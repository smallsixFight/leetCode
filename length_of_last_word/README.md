### 最后一个单词的长度（Length of Last Word）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/length-of-last-word/)

#### 题目
给定一个仅包含大小写字母和空格 ' ' 的字符串 s，返回其最后一个单词的长度。

如果字符串从左向右滚动显示，那么最后一个单词就是最后出现的单词。

如果不存在最后一个单词，请返回 0 。

说明：  
一个单词是指仅由字母组成、不包含任何空格的 最大子字符串。

示例:
```
输入: "Hello World"
输出: 5
```

#### 思路
从后面开始遍历，当 count 不为 0 并且当前值为 space 时直接返回；如果当前值不为 space 时 count +1。

#### 题外话
刷完上一题后看到这一题，感觉太简单了，都怀疑是不是看错题目了。

#### 伪代码
```
LENGTH-OF-LAST-WORD(s)
    count = 0
    for i = s.length -1 to 0
        if count != 0 && s[i] == ' '
            break
        if s[i] != ' '
            count += 1
    return count
```