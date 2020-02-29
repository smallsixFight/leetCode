### 电话号码的字母组合（Letter Combinations of a Phone Number）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/)

#### 题目
给定一个仅包含数字`2-9`的字符串，返回所有它能表示的字母组合。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。  
![17_telephone_keypad.png](https://i.loli.net/2020/02/29/4sxnp2LOZwDtj9A.png)

示例：
```
输入："23"
输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
```

#### 思路
这道题一开始没读懂，所以只能通过试运行一些例子去理解。

这道题有下面一些限制：

- 结果集的字符串长度与输入的字符串长度相等，即不包含子集；
- 如果输入的字符串包含`1`或`0`，那么不符合要求，直接返回空字符串数组；

解决思路如下：

- 使用一个 hash table 保存除`1`以外的数字跟字符的对应关系；
- 使用数组 res 保存结果集；遍历每一个字符，将新的结果集添加到 res，丢弃 res 原先的值；
- 如果判断字符为`0`或`1`，直接返回空字符串数组。

#### 伪代码
```
LETTER-COMBINATIONS(digits)
    dict = ['2': {'a', 'b', 'c'}, '3': {'d', 'e', 'f'}, '4': {'g', 'h', 'i'}, '5': {'j', 'k', 'l'}, 
                '6': {'m', 'n', 'o'}, '7': {'p', 'q', 'r', 's'}, '8': {'t', 'u', 'v'}, '9': {'w', 'x', 'y', 'z'}]
    res = new string Array()
    if digits.length < 1 || digits[0] == '0' || digits[0] == '1'
        return res
    for i = 0 to dict(digits[0]).length -1 // 初始化
        res.push(dict[digits[0]])
    for i = 1 to digits.length -1
        if digits[i] == '1'
            return empty string array
        arr = dict[digits[i]]
        l = res.length
        for j = 0 to l -1
            for v in arr
                res.push(res[j] + v)
        res = res[l:]
    return res
```