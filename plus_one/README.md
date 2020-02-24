### 加一（Plus One）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/plus-one/)

#### 题目
给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。

示例 1：
```
输入: [1,2,3]
输出: [1,2,4]
解释: 输入数组表示数字 123。
```
示例 2：
```
输入: [4,3,2,1]
输出: [4,3,2,2]
解释: 输入数组表示数字 4321。
```

#### 思路
这道题需要考虑的问题地方就是进位，如果存在进位则进行下一位的加一计算，没有则直接范围。  
唯一的特殊情况就是数组第一位进位的话需要创建一个新的数组来返回。

#### 伪代码
```
PLUS-ONE(digits)
    i = digits.length -1
    for i to 0
        digits[i] += 1
        if digits[i] < 10
            break
        digits[i] = 0
        i -= 1
    if i == -1 && digits[i+1] == 0
        arr = new array()
        arr[0] = 1
        for j = 0 to digits.length -1
            arr[j+1] = digits[j]
        return arr
    else
        return digits
```