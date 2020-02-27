### Z 字形变换（ZigZag Conversion）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/zigzag-conversion/)

#### 题目
将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为`"LEETCODEISHIRING"`行数为 3 时，排列如下：
```
L   C   I   R
E T O E S I I G
E   D   H   N
```
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如：`"LCIRETOESIIGEDHN"`。

请你实现这个将字符串进行指定行数变换的函数：
```
string convert(string s, int numRows);
```

示例 1：
```
输入: s = "LEETCODEISHIRING", numRows = 3
输出: "LCIRETOESIIGEDHN"
```
示例 2：
```
输入: s = "LEETCODEISHIRING", numRows = 4
输出: "LDREOEIIECIHNTSG"
解释:

L     D     R
E   O E   I I
E C   I H   N
T     S     G
```

#### 思路
可以发现如下规律：

- 竖轴上的`nextVal = s[last + (numRows-1)*2]`；
- 在`last %(numRows-1) != 0`时，nextVal 与当前值中间存在一个斜轴值——`val = s[numRows - (last %(numRows-1)) *2]`；

以上面为基础，循环执行并对临界值进行相应的处理即可。

#### 伪代码
CONVERT(s, numRows)
    divVal = numRows -1
    res = new array()
    for i = 0 to divVal
        p = i
        while p < s.length
            arr.add(s[p])
            if p % divVal != 0 && (divVal - p % divVal) * 2 + p < s.length
                arr.add(s[(divVal - p % divVal) * 2 + p])
            p += divVal *2
    return res.toString()