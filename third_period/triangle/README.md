### 三角形最小路径和（Triangle）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/triangle/)

#### 题目

给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。

例如，给定三角形：
```
[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
```
自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。

说明：

如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。

#### 思路
这道题可以发现最后一行的每个节点最小和取决于在第 n-1 行的邻结点的最小和，那么可以用 DP 来解决这个问题。

然后，可以发现，所谓三角形的总行数其实就是三角形最后行的总列数。

#### 伪代码
```
MINIMUM-TOTAL(triangle)
    if triangle == nil || triangle[0] == nil
        return 0
    dp = new Array()
    dp[0] = triangle[0][0]
    for i = 1 to triangle.length -1
        for j = triangle[i].length -1 to 0
            if j == triangle[i].length -1
                dp[j] = dp[j-1] + triangle[i][j]
            else
                if j == 0 || dp[j-1] > dp[j]
                    dp[j] += triangle[i][j]
                else
                    dp[j] = dp[j-1] + triangle[i][j]
    return min(dp)
```