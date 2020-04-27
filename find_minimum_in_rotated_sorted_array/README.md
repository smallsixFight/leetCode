###  寻找旋转排序数组中的最小值（Find Minimum in Rotated Sorted Array）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/)

#### 题目
假设按照升序排序的数组在预先未知的某个点上进行了旋转。

（例如，数组`[0,1,2,4,5,6,7]`可能变为`[4,5,6,7,0,1,2]`）。

请找出其中最小的元素。

你可以假设数组中不存在重复元素。

示例 1：
```
输入: [3,4,5,1,2]
输出: 1
```
示例 2：
```
输入: [4,5,6,7,0,1,2]
输出: 0
```

#### 思路
这道题没什么好说的，之前已经做过两个基于这道题的扩展题目。

一是从头遍历，返回第一个小于`nums[0]`的数；
二是使用二分法，找到`nums[n]<nums[0] && nums[n-1] >= nums[0]`的位置即可，也是前面扩展的题目要求，这里我就直接用二分法来做了。

#### 伪代码
```
FIND-MIN(nums)
    if nums.length == 0
        return 0
    low, high = 0, nums.length -1
    for low <= high
        mid = (low+high) >> 1
        if nums[mid] < nums[0] && nums[mid-1] >= nums[0]
            return nums[mid]
        elseif nums[mid] < nums[0]
            high = mid-1
        else
            low = mid+1
    return nums[0]
```