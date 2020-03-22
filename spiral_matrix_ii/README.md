### 螺旋矩阵 II（Spiral Matrix II）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/spiral-matrix-ii/)

#### 题目
给定一个正整数 n，生成一个包含 1 到 n^2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。

示例：
```
输入: 3
输出:
[
 [ 1, 2, 3 ],
 [ 8, 9, 4 ],
 [ 7, 6, 5 ]
]
```

#### 思路
这道题也可以用跟`螺旋矩阵`这道题一样的思路去做，之前是读数，现在是填数而已。

先把外层填充完，再去填充内层。

#### 伪代码
```
GENERATE-MATRIX(n)
    if n == 0
        return nil
    res = new Array()
    k = 0
    val = 0
    while k < (n+1) /2
        p = n -k -1
        for i = k to p
            res[k][i] = val +1
        for j = k +1 to p -1
            res[j][p] = val +1
        if p != k
            for i = p to k
                res[p][i] = val +1
            for j = p -1 to k +1
                res[k][j] = val +1
        k ++
    return res
```