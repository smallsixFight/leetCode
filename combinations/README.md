### 组合（Combinations）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/combinations/)

#### 题目
给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

示例：
```
输入: n = 4, k = 2
输出:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
```

#### 思路
从数学的角度来看，这就是一个`C(n k)`的组合问题，主要是程序如何实现并记录结果集。

这题可以用回溯法来做，先将每个可以作为起始节点的数字存到列表中，然后分别向后遍历获取符合要求的下一个值。

如果`n==k`，只有一个结果，直接返回即可。

#### 伪代码
```
COMBINE(n, k)
    res = new Array()
    if n == 0 || k == 0 || n < k
        res.push([]int{})
    if n == k
        arr = new Array()
        for i = 0 to n -1
            res.push(arr)
        return res
    for i = 1 to n -k +1
        res.push([]int{i})
    while res[0].length < k
        l = res.length
        for i = 0 to l -1
            val = res[i][res[i].length -1]
            j = 1
            while (k - res[i].length) <= (n -val -j +1)
                temp = res[i]
                temp.push(val+j)
                res.push(temp)
                j ++
        res = res[l:]
    return res
```