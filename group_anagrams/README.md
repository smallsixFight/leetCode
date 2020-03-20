### 字母异位词分组（Group Anagrams）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/group-anagrams/)

#### 题目
给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。

示例：
```
输入: ["eat", "tea", "tan", "ate", "nat", "bat"],
输出:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
```

**说明**

- 所有输入均为小写字母。
- 不考虑答案输出的顺序。

#### 思路一（超时）
这道题可以用 BFS 的思路来做，先把第一个单词入栈，并把该单词的中的字母分别存到一个 hash table 中和一个结果集，然后遍历剩下的单词，如果存在不同的字母，将该单词存到一个`next`的数组中，遍历结束后，从`next`取出一个单词，继续遍历，直到`next`数组为空。

#### 思路二
这个思路来自官方评论中的一个解法，使用 26 个质数表示 26 个字母，分别计算每个单词的乘积，相同的结果的单词存放到同一个结果集中。

这解法真的厉害，这就是大佬吧。

#### 思路三
这个思路是看其他提交记录学到的，耗时比思路二大了一倍，但能学多点就得多学点。

该算法使用的官方题解的第一个方法——排序数组分类，对每个单词进行字母排序，把排序相同的字符串存储起来，最终保存到结果集并返回。

#### 伪代码
```
GROUP-ANAGRAMS-TWO(strs)
    prime := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103}
    res = new Array()
    record = map[int]int
    count = 0
    for idx = 0 to strs.length -1
        t = 1
        for idx2 = strs[idx]
            t *= prime[str[idx][idx2] - 'a']
        if _, ok = record[t]; !ok
            record[t] = count
            count ++
            res.push(new Array())
        res[record[t]].push(strs[idx])
    return res

GROUP-ANAGRAMS(strs)
    next = strs
    res = new Array()
    for next.length > 0
        r = new Array()
        val = next[0]
        r.push(val)
        tempNext = new Array()
        m = new Map()
        if val != ""
            for i = 0 to val.length -1
                map[val[i]] ++
        for i = 1 to next.length -1
            if val.length != next[i].length
                tempNext.push(next[i])
				continue
            if next[i] == ""
                if val == ""
                    r.push(next[i])
                else
                    tempNext.push(next[i])
                continue
            else if val == ""
                tempNext.push(next[i])
				continue
            else
                mm = new Map()
                copy(mm, m)
                flag = true
                for j = 0 to next[i].length -1
                    if _, ok = mm[next[i][j]]; ok && mm[next[i][j]] > 0
                        mm[next[i][j]] --
                    else
                        tempNext.push(next[i])
                        flag = false
                        break
                if flag
                    r.push(next[i])
        next = tempNext
        res.push(r)
    return res
```