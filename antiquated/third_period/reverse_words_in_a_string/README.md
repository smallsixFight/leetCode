###  翻转字符串里的单词（Reverse Words in a String）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/reverse-words-in-a-string/)

#### 题目
给定一个字符串，逐个翻转字符串中的每个单词。

示例 1：
```
输入: "the sky is blue"
输出: "blue is sky the"
```
示例 2：
```
输入: "  hello world!  "
输出: "world! hello"
解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
```
示例 3：
```
输入: "a good   example"
输出: "example good a"
解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
```

说明：

- 无空格字符构成一个单词。
- 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
- 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。

进阶：

请选用 C 语言的用户尝试使用 O(1) 额外空间复杂度的原地解法。

#### 思路
使用双指针，从后向前遍历，将符合条件的字符串加入结果集，然后继续遍历直至结束。

#### 伪代码
```
REVERSE-WORDS(s)
    if s.length == 0
        return s
    res = new Array()
    p1, p2 = s.length -1, s.length -1
    for i = s.length - 1 to 0
        if s[p2] == ' '
            p2 --
        elseif s[p1] == ' '
            res.push(s[p1+1:p2+1] + ' ')
            p2 = p1
        p1 --
    if res.length > 0
        return res[:res.length-1].toString()
    return res.toString()
```