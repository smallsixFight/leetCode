### 解码方法（Decode Ways）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/decode-ways/)

#### 题目
一条包含字母`A-Z`的消息通过以下方式进行了编码：
```
'A' -> 1
'B' -> 2
...
'Z' -> 26
```

给定一个只包含数字的非空字符串，请计算解码方法的总数。

示例 1：
```
输入: "12"
输出: 2
解释: 它可以解码为 "AB"（1 2）或者 "L"（12）。
```
示例 2：
```
输入: "226"
输出: 3
解释: 它可以解码为 "BZ" (2 26), "VF" (22 6), 或者 "BBF" (2 2 6) 。
```

#### 思路一（不符合题目要求
手动画一下`12231421`和`12271421`的结果，可以发现一个规律：当`s[n]`能与`s[n-1]`组合成符合要求的数时，`f(n)=f(n-1)+f(n-2)`，当`s[n-1]s[n]>'26'`时，`f(n)=f(n-1)`。

这道题就可以用 DP 来做了。

**边界值**  
`0`是不能单独解码的，且只能在`1`和`2`后面。所以上面的条件需要进一步处理。

- 当`s[n]=0 && (s[n-1]<1 || s[n-1]>2`，f(n) = 0；
- 当`s[n]=0 && (s[n-1]=1 || s[n-1]==2)`，f(n) = f(n-2)。

#### 思路二
思路一是我一开始的想法，然后在写完代码在官网进行随便输测试案例的时候发现，`0`只能跟前一位进行组合，真 TM 坑人，我还是不明白`01`为什么不能被解析为`A`，题目并没有这个硬性要求啊。

所以思路二只是在符合题目要求的情况下对思路一进行调整，大致调整如下：

- 当`s[n]=0`时，如果`s[n-1]s[n] > 26 || s[n-1]s[n] < 1`，`f(n) = 0`；
- 当`s[n]=0 && (s[n-1]=1 || s[n-1]=2)`，f(n) = f(n-2)；
- 当`s[n-1] == 0 || s[n-1]s[n]>'26'`时，`f(n)=f(n-1)`；
- 当`s[n]`能与`s[n-1]`组合成符合要求的数时，`f(n)=f(n-1)+f(n-2)`。

#### 题外话
个人真感觉这道题的难点不在于 DP，而在于对 0 的处理。

另外，实际的代码`NumDecodingsTwo`没有使用数组，而是使用了三个局部变量，因为前面的并无必要保存。

#### 伪代码
```
NUM-DECODINGS-TWO(s)
    if s.length < 2
        if s.length == 1 && s[0] == '0'
            return 0
        return s.length
    res = new Array()
    if s[0] == '0'
        res.push(0)
        res.push(0)
    else
        if s[0] > '2' || (s[0] == '2' && s[1] > '6')
            res.push(1)
        else
            res.push(2)
    for i = 2 to s.length -1
        if s[i] == '0'
            if s[n-1] == '1' || s[n-1] == '2'
                res.push(res[n-2])
            else
                res.push(0)
        else
            if s[n-1] == '0' || s[n-1] > '2' || (s[n-1] == '2' && s[n] > '6')
                res.push(res[n-1])
            else
                res.push(res[n-1] + res[n=2])
        

NUM-DECODINGS(s)    // 不符合题目答案要求的错误答案
    res = new Array()
    if s[0] == '0'
        res.push(0)
        if s[1] == '0'
            res.push(0)
        else
            res.push(1)
    else
        res.push(1)
        if s[1] == '0'
            if s[0] > '2'
                res.push(0)
            else                // s[0] = '1' 或 s[0] = '2'
                res.push(1)
        else
            if (s[0] == '1') || (s[0] == '2' && s[1] < '7')
                res.push(2)
            else
                res.push(1)
    for i = 2 to s.length -1
        if s[i] == '0'
            if s[i-1]<'1' || s[i-1]>'2'
                res.push(0)
            elseif s[i-1]='1' || s[i-1]=='2'
                res.push(res[i-2])
        else
            if s[i-1] == '1' || (s[i-1] == '2' && s[i] < '7')
                res.push(res[i-1] + res[i-2])
            else
                res.push(res[i-1])
    return res
```