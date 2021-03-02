### 螺旋矩阵（Spiral Matrix）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/spiral-matrix/)

#### 题目
给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。

示例 1：
```
输入:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
输出: [1,2,3,6,9,8,7,4,5]
```
示例 2：
```
输入:
[
  [1, 2, 3, 4],
  [5, 6, 7, 8],
  [9,10,11,12]
]
输出: [1,2,3,4,8,12,11,10,9,5,6,7]
```

#### 思路
这道题可以用类似之前旋转图像那道题的方法来做，可以从先把外层读完，再读内层，依次类推。

#### 伪代码
```
SPIRAL-ORDER(matrix)
    if matrix.length == 0
        return nil
    res = new Array()
    m, n = matrix.length, matrix[0].length
    k = 0
    while k < (n+1) /2 && k < (m+1) /2
        p1, p2 = n -k -1, m -k -1
        for i = k to p1                 // 取顶行
            res.push(matrix[k][i])
        for j = k+1 to p2 -1            // 取右侧列
            res.push(matrix[j][p1])
        i = p1
        while i >= k && p2 != k         // 取底行，如果只有一行的情况下要避免重复
            res.push(matrix[p2][i])
            i --
        j = p2 -1
        while j > k && p1 != k          // 取左侧列，只有一列的情况下要避免重复
            res.push(matrix[j][k])
            j --
    return res
```