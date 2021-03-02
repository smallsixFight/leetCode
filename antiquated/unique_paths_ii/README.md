### 不同路径 II（Unique Paths II）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/unique-paths-ii/)

#### 题目
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。

现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？

![robot_maze.png](https://i.loli.net/2020/03/25/A1gvcCJmSwtVK9E.png)

网格中的障碍物和空位置分别用`1`和`0`来表示。

**说明：** m 和 n 的值均不超过 100。

示例：
```
输入:
[
  [0,0,0],
  [0,1,0],
  [0,0,0]
]
输出: 2
解释:
3x3 网格的正中间有一个障碍物。
从左上角到右下角一共有 2 条不同的路径：
1. 向右 -> 向右 -> 向下 -> 向下
2. 向下 -> 向下 -> 向右 -> 向右
```

#### 思路
这道题目依旧可以使用动态规划来做，只是遇到障碍物需要做一些处理而已。

#### 伪代码
```
UNIQUE-PATHS-WITH-OBSTACLES(obstacleGrid)
    per, cur = []int, []int
    for i = 0 to obstacleGrid.length -1
        for j = 0 to obstacleGrid[0].length -1
            if i == 0 && j == 0
                if obstacleGrid[0][0] != 1
                    cur[j] = 1
            elseif i == 0 && obstacleGrid[0][j] != 1
                cur[j] = cur[j-1]
            elseif j == 0 && obstacleGrid[i][0] != 1
                cur[j] = pre[j]
            else
                if obstacleGrid[i][j] != 1
                    cur[j] = cur[j-1] + pre[j]
        pre = cur
    return cur[j]
```