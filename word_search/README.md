### 单词搜索（Word Search）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/word-search/)

#### 题目
给定一个二维网格和一个单词，找出该单词是否存在于网格中。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

示例：
```
board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

给定 word = "ABCCED", 返回 true
给定 word = "SEE", 返回 true
给定 word = "ABCB", 返回 false
```

提示：

- `board`和`word`中只包含大写和小写英文字母。
- `1 <= board.length <= 200`
- `1 <= board[i].length <= 200`
- `1 <= word.length <= 10^3`

#### 思路
一开始想到 BFS，在写完伪代码后才发现虽然可以对节点进行回溯，但是 word 的回溯更是问题，代码也太长了，应该换一种方式，然后想到改用 DFS。

大致思路如下：

- 对每个符合要求的节点进行 DFS，如果存在符合要求的节点直接返回 true，不再继续探索，否则继续探索；全部节点符合条件的情况下仍然不符合条件，则直接返回 false。
- 由于题目限制同一个单元格的字符只能被使用一次，所以需要将当前使用过的单元格字符位置存起来。

#### 题外话一
一开始我使用了四个 if 分别向四个方向进行 DFS，遇到的一个问题就是如何让 record 回溯，一开始没想通用了个笨方法，每次都对原来的 record 进行 copy。

然后首次提交拿到 300+ms 的结果后，想了一下才恍然大悟，由于四次 DFS 是依次进行，并不是并行，那么只要前面的 DFS 将使用过节点标记复原就好了。

#### 题外话二
这次提交的代码在时间上只达到 30%+ 结果，有一个解法是在 word 嵌入`#`，不过目前没想明白，这是之后需要学习的一个算法。

其实之前做过的题目也提到这个算法，但我并没有去看，留着之后吧。

#### 伪代码
```
var words string
EXIST(board, word)
    if board.length == 0 || board[0].length == 0 || board.length * board[0].length < word.length
        return false
    words = word
    for i = 0 to board.length -1
        for j = 0 to board[0].length -1
            if DFS(board, 0, new Map(), i, j)
                return true
    return false
    
DFS(board, k, record, i, j)
    if i < 0 || i == board.length || j < 0 || j == board[0].length
        return false
    if record[[2]int{i, j}] || board[i][j] != words[k]
        return false
    record[[2]int{i, j}] = true
    if dfs(board, k+1, record, i-1, j) || dfs(board, k+1, record, i+1, j) || 
        dfs(board, k+1, record, i, j-1) || dfs(board, k+1, record, i j+1)
            return true
    record[[2]int{i, j}] = false
    return false
```