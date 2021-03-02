### 被围绕的区域（Surrounded Regions）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/surrounded-regions/)

#### 题目
给定一个二维的矩阵，包含`'X'`和`'O'`（字母 O）。

找到所有被`'X'`围绕的区域，并将这些区域里所有的`'O'`用`'X'`填充。

示例：
```
X X X X
X O O X
X X O X
X O X X
```
运行你的函数后，矩阵变为：
```
X X X X
X X X X
X X X X
X O X X
```
解释：

被围绕的区间不会存在于边界上，换句话说，任何边界上的`'O'`都不会被填充为`'X'`。任何不在边界上，或不与边界上的`'O'`相连的`'O'`最终都会被填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。

#### 思路
做一下思维的转换，从边界的`O`进行 DFS，将与边界连接的`O`转化为`#`，然后再次遍历，将`O`转换为`X`，`#`转换为`O`。

#### 伪代码
```
SOLVE(board)
    if board.length < 2 || board.[0] < 2
        return
    for i = 0 to board[0].length -1
        if board[0][i] == 'O'
            DFS(board, 0, i)
        if board[board.length-1][i] == 'O'
            DFS(board, board.length-1 , i)
    for i = 0 to board.length -1
        if board[i][0] == 'O'
            DFS(board, i, 0)
        if board[i][board[0].length-1] == 'O'
            DFS(board, i , board[0].length-1)
    for i = 0 to board.length -1
        for j = 0 to board[0].length -1
            if board[i][j] == '#'
                board[i][j] = 'O'
            elseif board[i][j] == 'O'
                board[i][j] = 'X'
                
DFS(board, i, j)
    board[i][j] = '#'
    if i > 0 && board[i-1][j] == 'O'
        DFS(board, i-1, j)
    if i < board.length -1 && board[i+1][j] == 'O'
        DFS(board, i +1, j)
    if j > 0 && board[i][j-1] == 'O'
        DFS(board, i, j-1)
    if j < board[0].length -1 && board[i][j+1] == 'O'
        DFS(board, i, j+1)
```