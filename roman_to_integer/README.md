### 罗马数字转整数（Roman To Integer）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/roman-to-integer/)

#### 题目
罗马数字包含以下七种字符：`I`，`V`，`X`，`L`，`C`，`D`和`M`。
```
字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
```
通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做`IIII`，而是`IV`。
数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为`IX`。这个特殊的规则只适用于以下六种情况：

- `I`可以放在`V`(5) 和`X`(10) 的左边，来表示 4 和 9。
- `X`可以放在`L`(50) 和`C`(100) 的左边，来表示 40 和 90。 
- `C`可以放在`D`(500) 和`M`(1000) 的左边，来表示 400 和 900。

给定一个罗马数字，将其转换成整数。输入确保在 1 到 3999 的范围内。

#### 思路
这题可以说是找规律，然后进行归纳，最后转换为代码。这题的规律有如下：

- 罗马数字小的通常在大的右边，除了 4、9、40、90、400 和 900；
- I 只能放在 V 和 X 的左边；
- X 只能放在 L 和 C 的左边；
- C 只能放在 D 和 M 的左边。

将上面的转换一下，在读到 I、X、C 时判断下一位是否更大，是的话就组合并加入 result，否则直接加入 result。

Solution2 是 LeetCode 上的另一种方式，根据判断对数字进行加和减，是一种更好的方式，可阅读性也好很多。

#### 伪代码
```
ROMAN-TO-INTEGER(s)
    result = 0
    for idx to s.length -1
        c = s[idx]
        if c == 'I'
            if idx +1 < s.length && s[i+1] == 'V'
                result += 4
                idx += 1
            elseif idx +1 < s.length && s[idx +1] == 'X'
                result += 9
                idx += 1
            else
                result += 1
        elseif c == 'X'
            if idx +1 < s.length && s[idx +1] == 'L'
                result += 40
                idx += 1
            elseif idx +1 < s.length && s[idx +1] == 'C'
                result += 90
                idx += 1
            else
                result += 10
        elseif c == 'C'
            if idx +1 < s.length && s[idx +1] == 'D'
                result += 400
                idx += 1
            elseif idx +1 < s.length && s[idx +1] == 'M'
                result += 900
                idx += 1
            else
                result += 100
        elseif c == 'V'
            result += 5
        elseif c == 'L'
            result += 50
        elseif c == 'D'
            result += 500
        else result += 1000
        idx += 1
    return result
```