### Excel 表列名称（Excel Sheet Column Title）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/excel-sheet-column-title/)

#### 题目
给定一个正整数，返回它在 Excel 表中相对应的列名称。

例如，
```
    1 -> A
    2 -> B
    3 -> C
    ...
    26 -> Z
    27 -> AA
    28 -> AB 
    ...
```

示例 1：
```
输入: 1
输出: "A"
```
示例 2：
```
输入: 28
输出: "AB"
```
示例 3：
```
`输入: 701
 输出: "ZY"
```

#### 思路
这道题当作一道 26 进制的题目来处理。

#### 伪代码
```
CONVERT-TO-TITLE(n)
    arr = new Array()
    while n > 0
        v = n %26
        if v == 0   // 如果是整数倍，那么则为 Z 字符
            v = 26
            n -= 1  // 进位需要减去 1
        arr.add(v -1 + 'A')
        n = n /26
    p1 ,p2 = 0, arr.length -1
    while p1 < p2
        arr[p1], arr[p2] = arr[p2], arr[p1]
    return arr.toString()
```