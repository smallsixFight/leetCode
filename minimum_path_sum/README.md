### 最小路径和（Minimum Path Sum）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/minimum-path-sum/)

#### 题目
给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

**说明：** 每次只能向下或者向右移动一步。

示例：
```
输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 7
解释: 因为路径 1→3→1→1→1 的总和最小。
```

#### 思路
这道题依旧可以使用 DP 方法来解决。

可以发现，到达当前位置的最短路径取决于`min(leftNum, topNUm)+grid[i][j]`。

另外，不需要额外的空间，只需要在原先的数组上面修改即可。

#### 伪代码
```
MIN-PATH-SUM(grid)
    if grid.length == 0 || grid[0].length == 0
        return 0
    for i = 0 to grid.length -1
        for j = 0 to grid.length -1
            if i == 0 && j == 0
                continue
            elseif i == 0
                grid[i][j] += grid[i][j-1]
            elseif j == 0
                grid[i][j] += grid[i-1][j]
            else
                if grid[i][j-1] < grid[i-1][j]
                    grid[i][j] += grid[i][j-1]
                else
                    grid[i][j] += grid[i-1][j]
    return grid[gird.length-1][grid[0].length-1]
```