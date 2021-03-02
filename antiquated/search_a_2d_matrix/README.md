### 搜索二维矩阵（Search a 2D Matrix）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/search-a-2d-matrix/)

#### 题目
编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：

- 每行中的整数从左到右按升序排列。
- 每行的第一个整数大于前一行的最后一个整数。

示例 1：
```
输入:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 3
输出: true
```
示例 2：
```
输入:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 13
输出: false
```

#### 思路
看到题目第一个想法就是二分查找，如何在矩阵实现二分查找才是需要思考的地方。

有鉴于前面的经验，可以做两次二分查找，第一次找到行，第二次在该行进行查找。

#### 伪代码
```
SEARCH-MATRIX(matrix, target)
    if matrix.length == 0 || matrix[0] == 0
        return false
    m, n = matrix.length, matrix[0].length
    if target < matrix[0][0] || target > matrix[m-1][n-1]
        return false
    low, high = 0, m -1
    for low <= high
        mid = (low + high) /2
        if target == matrix[mid][0] || target == matrix[mid][n-1]
            return true
        elseif target > matrix[mid][n-1]
            low = mid +1
        else target < matrix[mid][0]
            high = mid -1
    row = low
    low, high = 0, n-1
    for low <= high
        mid = (low + high) /2
        if matrix[row][mid] == target
            return true
        elseif matrix[row][mid] < target
            low = mid +1
        else
            high = mid -1
    return false
```