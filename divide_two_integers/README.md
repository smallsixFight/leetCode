### 两数相除（Divide Two Integers）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/divide-two-integers/)

#### 题目
给定两个整数，被除数`dividend`和除数`divisor`。将两数相除，要求不使用乘法、除法和`mod`运算符。

返回被除数`dividend`除以除数`divisor`得到的商。

示例 1：
```
输入: dividend = 10, divisor = 3
输出: 3
```

示例 2：
```
输入: dividend = 7, divisor = -3
输出: -2
```

说明：

- 被除数和除数均为 32 位有符号整数。
- 除数不为 0。
0 假设我们的环境只能存储 32 位有符号整数，其数值范围是`[−2^31,  2^31 − 1]`。本题中，如果除法结果溢出，则返回`2^31 − 1`。

#### 思路
这题可以用左移右移和递归来完成。

#### 伪代码
```
DIVIDE(dividend, divisor)
    flag = false
    if dividend < 0
        dividend = -dividend
        flag = !flag
    if divisor < 0
        divisor = -divisor
        flag = !flag
    if dividend < divisor
        return 0
    res = 1
    temp = divisor << 1
    for temp < dividend
        temp <<= 1
        res <<= 1
    temp >> 1
    if temp < dividend
        res += DIVIDE(dividend -temp, divisor)
    if flag
        res = -res
    if res > math.MaxInt32
        res = math.MaxInt32
    return res
```
