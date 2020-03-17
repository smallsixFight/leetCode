### 组合总和（Combination Sum）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/combination-sum/)

#### 题目
给定一个**无重复元素**的数组`candidates`和一个目标数`target`，找出`candidates`中所有可以使数字和为`target`的组合。

`candidates`中的数字可以无限制重复被选取。

**说明**

- 所有数字（包括`target`）都是正整数。
- 解集不能包含重复的组合。

示例 1：
```
输入: candidates = [2,3,6,7], target = 7,
所求解集为:
[
  [7],
  [2,2,3]
]
```
示例 2：
```
输入: candidates = [2,3,5], target = 8,
所求解集为:
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]
```

#### 思路
这道题可以使用**回溯+剪枝**来解决，大致思路如下：

- 如果当前值使得`sum == target`，那么添加到结果集；由于升序和元素不重复，那么可直接回溯；
- 如果当前值使得`sum > target`，直接回溯；
- 如果当前值使得`sum < target`，由于元素可以重复使用，那么以这个元素作为起点，继续向后探索。

#### 伪代码
```
res = new Array()
var t int
c = new Array()
COMBINATION-SUM(candidates, target)
    if candidates.length == 0
        return res
    t = target
    c = candidates
    for i = 0 to c.length -1
        DEATH-EXPLORE(0, i, new Array())
    return res

DEATH-EXPLORE(sum, i, arr)
    sum += c[i]
    if sum > t  // 剪枝
        return
    if sum == t
        temp = arr
        temp.push(c[i])
        res.push(temp)
        return
    arr.push(c[i])
    for j = i to c.length -1
        DEATH-EXPLORE(sum, j, arr)
```