### 全排列 II（Permutations II）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/permutations-ii/)

#### 题目
给定一个可包含重复数字的序列，返回所有不重复的全排列。

示例：
```
输入: [1,1,2]
输出:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]
```

#### 思路
这道题跟上道题相比，就多了一个去重的步骤，相比前一题就加了两行代码。

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
                if res[j][x] == res[j][res[j].[length -1]]
                    break
                temp = res[j]
                t = temp[x]
                temp[x] = temp[temp.length -1]
                temp[temp.length - 1] = temp[x]
                res.push(temp)
    return res
```

