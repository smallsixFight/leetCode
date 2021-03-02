### 杨辉三角 II（Pascal's Triangle II）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/pascals-triangle-ii/)

#### 题目
给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。

![PascalTriangleAnimated2.gif](https://i.loli.net/2020/04/09/NGarly6DjJ1tpxA.gif)

在杨辉三角中，每个数是它左上方和右上方的数的和。

示例：
```
输入: 3
输出: [1,3,3,1]
```

进阶：  
你可以优化你的算法到 O(k) 空间复杂度吗？

#### 思路
可以在前一题的基础上稍微修改一下，减少一下空间的使用。因为可以发现，第`n`行只取决与第`n-1`行的数组，跟之前的结果集不会产生直接联系，那么只需要根据前一个结果集来修改就好了。

这个解法挺简单的，可能是因为之前做过中等类型的这类题目吧。

#### 吐槽
这道题的开头直接坑了我一下，k 是索引而不是行数，还有没有直接提交。

#### 伪代码
```
GET-ROW(rowIndex)
    if rowIndex == 0
        return nil
    elseif rowIndex == 1
        return []int{1}
    res = []int{1, 1}
    for i = 3 to rowIndex
        res[rowIndex-1] = 1
        for k = rowIndex -2 to 1
            res[k] += res[k-1]
    return res
```