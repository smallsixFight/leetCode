### 组合总和（Combination Sum II）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/combination-sum-ii/)

#### 题目
给定一个数组`candidates`和一个目标数`target`，找出`candidates`中所有可以使数字和为`target`的组合。

`candidates`中的每个数字在每个组合中只能使用一次。

**说明**

- 所有数字（包括目标数）都是正整数。
- 解集不能包含重复的组合。

示例 1：
```
输入: candidates = [10,1,2,7,6,1,5], target = 8,
所求解集为:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
```
示例 2：
```
输入: candidates = [2,5,2,1,2], target = 5,
所求解集为:
[
  [1,2,2],
  [5]
]
```

#### 思路
这道题跟前一道题的区别是需要去重以及同一个位置的元素不能重复出现。

去重这个一开始看着有些眼熟，想了想发现跟之前没做出来的`三数之和`类似，可以结合`三数之和`的去重方式来实现，大致思路如下：

- 首先对数组`candidates`进行排序；
- 由于是对每个节点进行 DFS，可以发现，如果当前节点同前一个节点相同，那么它的解一定是前一个结点的子集，那么就可以跳过不进行 DFS，也达到了去重的目的。

#### 伪代码
```
COMBINATION-SUM(candidates, target)
    res = new Array()
    sort(candidates)
    for i := 0 to candidates.length -1 && candidates[i] <= target
        if i > 0 && candidates[i] == candidates[i-1]
            continue
        DEATH-EXPLORE(i, target, new Array())
    return res
        
    
DEATH-EXPLORE(i, target, arr)
    if candidates[i] > target
        return
    if candidates[i] == target
        temp = arr
        temp.push(candidates[i])
        res.push(temp)
        return
    target -= candidates[i]
    arr.push(candidates[i])
    for j = i +1 to candidates.length -1
        DEATH-EXPLORE(j, target, arr)
        while j +1 < candidates.length && candidates[j+1] == candidates[j]
            j ++
```