### 岛屿数量（Number of Islands）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/number-of-islands/)

#### 题目
给你一个由`'1'`（陆地）和`'0'`（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。

示例 1：
```
输入:
11110
11010
11000
00000
输出: 1
```
示例 2：
```
输入:
11000
11000
00100
00011
输出: 3
解释: 每座岛屿只能由水平和/或竖直方向上相邻的陆地连接而成。
```

#### 思路
遍历整个二维网格，如果遇到`1`，则岛屿数量 +1，然后以对该点进行 BFS，将所有`1`改为`-1`。

#### 伪代码
```
NUM-ISLANDS(grid)
    res = 0
    for i = 0 to grid.length -1
        for j = 0 to grid[i].length -1
            if grid[i][j] == 1
                grid[i][j] = -1
                BFS(i, j)
    return res

BFS(x, y)
    queue = new Queue()
    queue.enqueue([]int{x, y})
    for !queue.isEmpty()
        l = queue.length
        for i = 0 to l -1
            position = queue.dequeue()
            a, b = position[0], position[1]
            grid[a][b] = -1
            if a > 0 && grid[a-1][b] == 1
                grid[a-1][b] = -1
                queue.enqueue(grid[a-1][b])
            if a < grid[a].length && grid[a+1][b] == 1
                grid[a+1][b] = -1
                queue.enqueue(grid[a+1][b])
            if b > 0 && grid[a][b-1] == 1
                grid[a][b-1] = -1
                queue.enqueue(grid[a][b-1])
            if b < grid.length && grid[a][b+1] == 1
                grid[a][b+1] = -1
                queue.enqueue(grid[a][b+1])
```