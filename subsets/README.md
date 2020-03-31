### 子集（Subsets）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/subsets/)

#### 题目
给定一组**不含重复元素**的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例：
```
输入: nums = [1,2,3]
输出:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
```

#### 思路
这题可以用 DFS 的方式来做。

#### 伪代码
```
res = new Array()
SUBSETS(nums)
    res.push([]int{})
    for i = 0 to nums.length -1
        res.push([]int{i})
    for i = 1 to nums.length -1
        DFS(res(i), i)
        
DFS(arr, s)
    if s == nums.length
        return
    for i = s to nums.length -1
        temp = arr
        temp.push(nums[i])
        res.push(temp)
        DFS(temp, i+1)
```