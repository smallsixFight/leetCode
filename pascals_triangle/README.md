### 杨辉三角（Pascal's Triangle）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/pascals-triangle/)

#### 题目
给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。

![PascalTriangleAnimated2.gif](https://i.loli.net/2020/04/09/NGarly6DjJ1tpxA.gif)

在杨辉三角中，每个数是它左上方和右上方的数的和。

示例：
```
输入: 5
输出:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]
```

#### 思路
这道题可以用 DP 来做（不用自己画图真好）。

#### 伪代码
```
GENERATE(numRows)
    if numRows == 0
        return nil
    elseif numRows == 1
        return [][]int{{1}}
    elseif numRows == 2
        return [][]int{{1}, {1, 1}}
    res = new Array()
    res.push([]int{1}, []int{1, 1})
    for i = 3 to numRows
        arr = new Array()
        for k = 0 to i -1
            if k == 0 || k == i -1
                arr.push(1)
            else
                arr.push(res[i-2][k-1] + res[i-2][k])
        res.push(arr)
    return res
```
    

