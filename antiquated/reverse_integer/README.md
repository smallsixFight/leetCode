### 整数反转（Reverse Integer）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/reverse-integer/)

#### 题目
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

**注意:**  
假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为[−2^31,  2^31 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。

示例 1：
```
输入: 123
输出: 321
```
示例 2：
```
输入: -123
输出: -321
```
示例 3：
```
输入: 120
输出: 21
```

#### 思路
1. 对整数取模可以得到个位上的值，整数除 10 会向下取整；以此类推获取所有数字；
2. 将获取的数字组装成一个新的整数：result = result * 10 + newInt32；
3. 针对溢出：正数溢出存在 result > (2^31)/10 或 
(result = (2^31)/10 && newInt32 > 7) 两种情况；
负数存在 result > (2^31)/10 或 (result = (2^31) && newInt32 > -8) 两种导致溢出的情况。

#### 伪代码
```
REVERSE(x)
    result = 0
    if x > 0
        limitVal = maxInt32 /10
    else limitVal = minInt32 /10
    while x != 0
        val = x %10
        x /= 10
        if (result > limitVal) || (result == limitVal && val > 7)
            return 0
        if (result < limitVal) || (result == limitVal && val < -8)
            return 0
        result = result *10 +val
    return result
```