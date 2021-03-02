### 在排序数组中查找元素的第一个和最后一个位置（Find First and Last Position of Element in Sorted Array）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/)

#### 题目
给定一个按照升序排列的整数数组`nums`，和一个目标值`target`。找出给定目标值在数组中的开始位置和结束位置。

你的算法时间复杂度必须是 O(log n) 级别。

如果数组中不存在目标值，返回`[-1, -1]`。

示例 1：
```
输入: nums = [5,7,7,8,8,10], target = 8
输出: [3,4]
```
示例 2：
```
输入: nums = [5,7,7,8,8,10], target = 6
输出: [-1,-1]
```

#### 思路
这道题可以用跟上一道题——搜索旋转数组一样的解决方式，通过两次二分查找来解决。

第一次查找`nums[mid] == target and nums[mid-1] < target`，那么 mid 就是开始位置；

如果第一次能够查找到，则进行第二次查找，第二次查找`nums[mid] == target and nums[mid+1] > target`的位置，这个必定能够找到。

第二次查找可以用第一次查找到的开始位置作为起点。

#### 伪代码
```
SEARCH-RANGE(nums, target)
    res = [-1, -1]
    if nums.length == 0
        return res
    low, high = 0, nums.length -1
    for low <= high
        mid = (low + high) /2
        if nums[mid] == target
            if mid == 0 || (mid > 0 && nums[mid-1] != target)
                res[0] = mid
                break
            else
                high = mid -1
        elseif nums[mid] < target
            low = mid +1
        else
            high = mid -1
    if res[0] == -1
        return res
    if res[0] == nums.length -1
        res[1] = res[0]
        return res
    low, high = res[0], nums.length -1
    for low <= high
        mid = (low + high) /2
        if nums[mid] == target
            if mid == nums.length -1 || (mid +1 < nums.length && nums[mid+1] != target)
                res[1] = mid
                break
            else
                low = mid +1
        elseif nums[mid] < target
            low = mid +1
        else
            high = mid -1
    return res
```