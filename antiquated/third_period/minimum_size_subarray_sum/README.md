### 长度最小的子数组（Minimum Size Subarray Sum）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/minimum-size-subarray-sum/)

#### 题目
给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的连续子数组，并返回其长度。如果不存在符合条件的连续子数组，返回 0。

示例：
```
输入: s = 7, nums = [2,3,1,2,4,3]
输出: 2
解释: 子数组 [4,3] 是该条件下的长度最小的连续子数组。
```

进阶：

如果你已经完成了O(n) 时间复杂度的解法, 请尝试 O(n log n) 时间复杂度的解法。

#### 思路一
使用双指针法，从左向右遍历，先找到一个符合条件子数组。循环下面的操作：

- 左指针向右移动一位，如果仍然符合条件，则继续向右移动；
- 如果左指针移动后不符合条件，右指针向右移动，知道找到下一个符合条件的。

#### 伪代码
```
MIN-SUB-ARRAY-LEN(s, nums)
    if nums.length == 0
        return 0
    res, curSum = 0, nums[0]
    if curSum >= s
        return 1
    p1, p2 = 0, 1
    while p1 < nums.length && p2 <= nums.length
        if curSum >= s
            if p2 - p1 < res
                res = p2 - p1
            curSum -= nums[p1]
            p1 ++
        else
            if p2 == nums.length
                break
            curSum += nums[p2]
            p2 ++
    return res
```