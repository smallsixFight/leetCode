### 子集 II（Subsets II）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/subsets-ii/)

#### 题目
给定一个可能包含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例：
```
输入: [1,2,2]
输出:
[
  [2],
  [1],
  [1,2,2],
  [2,2],
  [1,2],
  []
]
```

#### 思路
这一题依旧可以使用 DFS 来处理，只是需要加多去重的步骤。

- 先给 nums 排序；
- 在初始化的时候，如果当前字符与前一个字符相同，不进行 DFS；
- 在 DFS 的时候，如果当前字符不是第一个且与之前的字符相同，不行 DFS。

#### 伪代码
```
res = new Array()
SUBSETS-WITH-DUP(nums)
    sort(nums)
    res.push([]int{})
    positionRecord = new Array()
    for i = 0 to nums.length -1
        if i > 0 && nums[i] == nums[i-1]
            continue
        res.push([]int{nums[i]})
        positionRecord.push(i+1)
    for i = 0 to positionRecord.length -1
        DFS(res[i+1], positionRecord(i))
        
DFS(arr, s)
    if s>= nums.length
        return
    for i = s to nums.length -1
        if i != s && nums[i] == nums[i-1]
            continue
        temp = arr
        temp.push(nums[i])
        res.push(temp)
        DFS(temp, i+1)
```