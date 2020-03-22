### 合并区间（Merge Intervals）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/merge-intervals/)

#### 题目
给出一个区间的集合，请合并所有重叠的区间。

示例 1：
```
输入: [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6]。
```
示例 2：
```
输入: [[1,4],[4,5]]
输出: [[1,5]]
解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
```

#### 思路
可以发现如果一个区间的末位置大于其他的区间的始位置，那么这两个区间重叠，可以合并。

由于题目没有明确区间是有序的，所以需要先按照区间的始位置进行排序，然后从头开始比较，直到遇到不重叠的区间，将重叠的区间合并并记录，从不重叠的区间进行下一轮比较合并。

#### 伪代码
```
MERGE(intervals)
    res = new Array()
    sort(intervals) // 按区间的始位置排序
    for i = 0 to intervals.length -1
        val = intervals[i]
        // 获取当前区间重叠最后的区间
        j = i +1
        while j < intervals.length && intervals[j][0] <= val[1]
            if intervals[j][1] > val[1]
                val[1] = intervals[j][1] 
            j += 1
        i = j
        res.push(val)
    return res
```