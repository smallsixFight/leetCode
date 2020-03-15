### 搜索旋转排序数组（Search in Rotated Sorted Array）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/search-in-rotated-sorted-array/)

#### 题目
假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组`[0,1,2,4,5,6,7]`可能变为`[4,5,6,7,0,1,2]`)。

搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。

你可以假设数组中不存在重复的元素。

你的算法时间复杂度必须是 O(log n) 级别。

示例 1：
```
输入: nums = [4,5,6,7,0,1,2], target = 0
输出: 4
```
示例 2：
```
输入: nums = [4,5,6,7,0,1,2], target = 3
输出: -1
```

#### 思路
因为使用二分法的前提是有序，旋转后其实也是有序的，只是位置变换了下而已，依旧可以二分查找，但是需先找到旋转点。

旋转后，旋转点前后的数组依旧是有序的，那么只需要找到比当前`nums[0]`大的`nums[i] > nums[0] and nums[i+i] < nums[0]`，那么`nums[i+1]`就是最小值的位置，也就是旋转点。

关于如何去找旋转点，其实也可以使用二分法来查找，找到后使用二分法来查找 target，所以这道题需要使用两次二分法。

#### 伪代码
```
SEARCH(nums, target)
    if nums.length  == 0
        return -1
    t1 = nums[0]
    low, high = 0, nums.length -1
    minIdx = 0      // 没旋转的话就是 0
    for low <= high
        mid = (low + high) /2
        if t1 <= nums[mid] && mid +1 < nums.length
            if t1 > nums[mid+1]
                midIdx = mid +1
                break
            else
                low = mid +1
        elseif t1 > nums[mid]
            high = mid -1
        else
            break
    if target < nums[0]
        low = minIdx
        high = nums.length -1
    else target > nums[0]
        low = 1
        if midIdx != 0
            high = minIdx -1
        else
            high = nums.length -1
    else
        return 0
    for low <= high
        mid = (low + high) /2
        if target == nums[mid]
            return mid
        elseif target < nums[mid]
            high = mid -1
        else
            low = mid +1
    return -1
```