### 解数独（Sudoku Solver）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/sudoku-solver/)

#### 题目
编写一个程序，通过已填充的空格来解决数独问题。

一个数独的解法需遵循如下规则：  
1. 数字`1-9`在每一行只能出现一次。
2. 数字`1-9`在每一列只能出现一次。
3. 数字`1-9`在每一个以粗实线分隔的`3x3`宫内只能出现一次。

空白格用`'.'`表示。

![base](http://upload.wikimedia.org/wikipedia/commons/thumb/f/ff/Sudoku-by-L2G-20050714.svg/250px-Sudoku-by-L2G-20050714.svg.png)

一个数独。

![250px-Sudoku-by-L2G-20050714_solution.svg.png](https://i.loli.net/2020/04/12/KsujbE6mPGnh1xD.png)

答案被标成红色。

提示：

- 给定的数独序列只包含数字`1-9`和字符`'.'`。
- 你可以假设给定的数独只有唯一解。
- 给定数独永远是`9x9`形式的。

#### 思路
我个人是基本没了解过数独的，然后去百度了一些相关的解法，其中就有`摒除法`，编程方面的实现来讲的话，就是使用`回溯法`。

从头开始遍历，将每个可能的数字放到表格中，然后继续放下一个数字。如果下一个单元格不能放入任何数字且没达到表格最后一位，那么回溯到前一个数字，尝试下一个可能的数字，不行则继续回溯，直到找到答案。

由于数字的表格长度和宽度都是 9，所以使用的空间复杂度也是常量，时间复杂度也是常量。

#### 伪代码
```
SOLVE-SUDOKU(board)
    rows, cols, boxes = [9][10]int, [9][10]int, [9][10]int
    for i = 0 to 8
        for j = 0 to 8
            if board[i][j] != '.'
                val = board[i][j] - '0'
                idx = (i / 3 ) * 3 + j / 3
                PLACE-NUMBER(val, i, j)
    BACKTRACK(0, 0)

BACKTRACK(row, col)
    if board[row][col] == '.'
        for i = 1 to 9
            if COULD-PLACE(i, row, col)
                PLACE-NUMBER(i, row, col)
                PLACE-NEXT-NUMBER(row, col)
            if !isSolve
                REMOVE(i, row, col)
    else
        PLACE-NEXT-NUMBER(row, col)

PLACE-NUMBER(num, row, col)
    idx = (row / 3 ) * 3 + col / 3
    rows[row][num] ++
    cols[col][num] ++
    boxes[idx][num] ++
    board[row][col] = num + '0'

PLACE-NEXT-NUMBER(row, col)
    if row == 9 && col == 9
        isSolve = true
    else
        if col == 8
            BACKTRACK(row+1, 0)
        else
            BACKTRACK(row, col+1)
    
REMOVE-NUMBER(num, row, col)
    idx = (row / 3 ) * 3 + col / 3
    rows[row][num] --
    cols[col][num] --
    boxes[idx][num] --
    board[row][col] = '.'

COULD-PLACE(num, row, col)
    idx = (row / 3 ) * 3 + col / 3
    return rows[row][num] + cols[col][num] + boxes[idx][num] == 0
```