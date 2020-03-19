### 全排列（Permutations）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/permutations/)

#### 题目
给定一个**没有重复**数字的序列，返回其所有可能的全排列。

示例：
```
输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
```

#### 思路
这道题可以用动态规划来完成，可以发现，第`n`个的结果集由第`n-1`个的结果集与第`n`个值进行组合，然后新的组合进行排列组合生成新的结果集。

#### 伪代码
```
PERMUTE(nums)
    res = [][]int
    res[0] = [nums[0]]
    for i = 1 to nums.length -1
        for j = 0 to res.length -1
            res[j].push(nums[i])
        resLen = res.length
        for j = 0 to resLen -1
            for x = 0 to res[j].length -2
                temp = res[j]
                t = temp[x]
                temp[x] = temp[temp.length -1]
                temp[temp.length - 1] = temp[x]
                res.push(temp)
    return res
```