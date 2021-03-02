### 复原 IP 地址（Restore IP Addresses）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/restore-ip-addresses/)

#### 题目
给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。

示例：
```
输入: "25525511135"
输出: ["255.255.11.135", "255.255.111.35"]
```

#### 思路
这道题可以用回溯法来做，然后来整理一下限制条件：

- 每个位置上长度最小为 1 位，最大为 3 位；
- 每个位置上的最大值不能超过 255。
- 长度大于 0 时，开头不能为 0。

#### 题外话
我一开始理解错题目了，以为可以灵活处理无效的 0，然而并不是，而是每个字符都必须用到，所以才有上面的第三条限制条件。

#### 伪代码
```
res = new Array()
RESTORE-IP-ADDRESSES(s)
    if s.length < 4 || s.length > 12
        return res
    BACKTRACK(s, 0, 4, "")
    return res

BACKTRACK(s, idx, count, sb)
    if s[idx:].length < count || s[idx:].length > 3 * count
        return
    count --
    for i = 1 to 3
        if s[idx+1:].length < count
            return
        if count > 0
            if i > 1 && ip[idx] == '0'
                continue
            if i == 3
                if ip[idx] > '2' || (ip[idx] == '2' && ip[idx+1] > '5') ||
                        (ip[idx] == '2' && ip[idx+1] == '5' && ip[idx+2] > '5')
                    return
            str = ip[idx:idx+i]
            while str[0] == '0' && str.length > 1
                str = str[1:]
            BACKTRACK(idx+i, count, sb + ".")
        else
            while ip[idx:].length > 1 && ip[idx] == '0'
                return
            if ip[idx:] == 3
                if ip[idx] > '2' || (ip[idx] == '2' && ip[idx+1] > '5') ||
                    (ip[idx] == '2' && ip[idx+1] == '5' && ip[idx+2] > '5')
                    continue
            res.push(sb + ip[idx:])
            break
```