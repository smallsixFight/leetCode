### 字符串相乘（Multiply Strings）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/multiply-strings/)

#### 题目
给定两个以字符串形式表示的非负整数`num1`和`num2`，返回`num1`和`num2`的乘积，它们的乘积也表示为字符串形式。

示例 1：
```
输入: num1 = "2", num2 = "3"
输出: "6"
```
示例 2：
```
输入: num1 = "123", num2 = "456"
输出: "56088"
```

**说明**

1. `num1`和`num2`的长度小于110。
2. `num1`和`num2`只包含数字 0-9。
3. `num1`和`num2`均不以零开头，除非是数字 0 本身。
4. 不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。

#### 思路
这道题让我想起小学乘法计算方式，不过需要考虑溢出的情况，所以不能直接相乘后进行相加。

#### 伪代码
```
MULTIPLY(num1, num2)
    if num1 == "0" || num2 == "0"
        return 0
    sum = new Array()
    for i = num1.length -1 to 0
        adv = 0
        for j = num2.length -1 to 0
            mul = (num1[i] - '0') * (num2[j] - '0') + adv
            adv = mul /10
            sum[i+j+1] += adv % 10
            if sum[i+j+1] > 9
                sum[i+j+1] = sum[i+j+1] %10
                adv ++
        if adv > 0
            sum[i] = adv
    if sum[0] == 0
        sum = sum[1:]
    for i = 0 to sum.length -1
        sum[i] += '0'
    return sum.toString()
```