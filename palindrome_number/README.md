### 回文数（Palindrome Number）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/palindrome-number/)

#### 题目
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例1：
```
输入: 121
输出: true
```
示例2：
```
输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
```
**进阶:**  
你能不将整数转为字符串来解决这个问题吗？

#### 思路
临界问题：由示例可以看出，负整数均不是回文数；另外，如果最后一个数字是 0,则第一位必须是 0,这样的话则有 0 符合该条件。
第一种方法是直接将数字转换为字符串，使用双指针进行匹配比较；
第二种方法是 LeetCode 官方的方法，对反转整数的其中一半数字，然后与另外一半进行大小比较。

#### 伪代码
```
FirstMethod(x)
    if x < 0 || (x %10 ==0 && x != 0)
        return false
    str = x.toString()
    i = 0
    j = str.length -1
    while i < j
        if str[i] != str[j]
            return false
    i ++
    j --
    return true
```
```
SecondMethod(x)
    if x < 0 || (x %10 ==0 && x != 0)
        return false
    reverseVal = 0
    while x > reverseVal
        reverseVal = reverseVal * 10 + x % 10
        x /= 10
    return reverseVal == x || reverseVal /10 == x
```