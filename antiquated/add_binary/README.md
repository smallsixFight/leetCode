### 二进制求和（Add Binary）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/add-binary/)

#### 题目
给定两个二进制字符串，返回他们的和（用二进制表示）。

输入为非空字符串且只包含数字 1 和 0。

示例 1：
```
输入: a = "11", b = "1"
输出: "100"
```
示例 2：
```
输入: a = "1010", b = "1011"
输出: "10101"
```

#### 思路
这道题算是`PlusOne`的简单扩展，思路上基本上一致。  
使用双指针从后向前遍历，需要考虑当前是否进位和后面是否进位。  
这道题我基于 ASCII 编码表来计算。

#### 伪代码
```
ADD-BINARY(a, b)
    l = math.Max(a.length, b.length)
    let arr[0...l] be new byte array
    i, j = a.length -1, b.length -1
    plusVal = 0
    for l to 0
        if i >= 0 && j >= 0
            temp = a[i] + a[j] + plusVal
        else if i >= 0
            temp = a[i] + plusVal
        else if j >= 0
            temp = a[j] + plusVal
        else 
            temp = plusVal
        if temp > 1
            plusVal = 1
        else plusVal = 0
        arr[l] = temp % 2
        i --
        j --
        l --
    return arr.toString()
``` 