### 括号生成（Generate Parentheses）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/generate-parentheses/)


#### 题目
给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且**有效的**括号组合。

例如，给出 n = 3，生成结果为：
```
[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
```

#### 思路
这道题可以回溯算法（我更喜欢叫“探索法”），其实也是说是深度优先探索，当目前探索结点的值符合要求，则继续向下探索，直到结尾将符合要求的结果加入结果集；一旦遇到不符合条件的则直接返回（回溯）到上一个结点，探索其它条件是否符合；以此类推。

那么，观察后可以发现，要符合**有效的**的括号组合，需要符合下面的条件：

- 如果探索的符号是`)`，那么必须要有一个`(`进行组合，否则必然是无效组合；
- 如果`（`的数量大于 n，则必然是无效组合；
- 如果探索结束，`(`总数大于零，则必然是无效的。

#### 伪代码
```
GENERATE-PARENTHESIS(n)
    res = new string Array()
    DEPTH-EXPLORE("(", 1, res, n)
    return res
    
DEPTH-EXPLORE(preStr, count, res, n)
    if preStr.length == 2n
        if count == 0
            res.push(preStr)
        return
    if count < n
        DEPTH-EXPLORE(preStr + '(', count +1, res, n)
    if count > 0
        DEPTH-EXPLORE(preStr + ')', count -1, res, n)
```