### 搜索旋转排序数组 II（Search in Rotated Sorted Array II）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii/)

#### 题目
假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组`[0,0,1,2,2,5,6]`可能变为`[2,5,6,0,0,1,2]`)。

编写一个函数来判断给定的目标值是否存在于数组中。若存在返回`true`，否则返回`false`。

示例 1：
```
输入: nums = [2,5,6,0,0,1,2], target = 0
输出: true
```
示例 2：
```
输入: nums = [2,5,6,0,0,1,2], target = 3
输出: false
```

进阶：

- 这是`搜索旋转排序数组`的延伸题目，本题中的`nums`可能包含重复元素。
- 这会影响到程序的时间复杂度吗？会有怎样的影响，为什么？

#### 思路
这道题跟`搜索旋转排序数组`同样可以使用两次二分查找，不过这题允许重复元素的存在，所以出现差异的地方是在第一次二分查找，第二次的二分查找是没有变化的。

在`搜索旋转排序数组`只需要两次二分查找，分别用于查找旋转点和搜索是否存在 target。然而在存在重复元素的本题中，举一个简单明了的例子——`[1, 1, 2, 1, 1, 1, 1, 1]`和`[1, 1, 1, 1, 1, 1, 2, 1]`，
会发现一次二分不能准确的找到 low 和 high，需要分别对左侧和右侧继续进行二分查找才行，这就需要递归进行二分查找了。

所以，关于上面说到的会影响时间复杂度的问题，答案是肯定的，一旦出现前面例子说的情况，那么就需要进行递归二分。

#### 伪代码
```
SEARCH(nums, target)
    if nums.length == 0
        return false
    if nums[0] == target || nums[nums.length-1] == target
        return true
    low, high = 0, nums.length -1
    p = FIND-ROTATE-POSITION(nums, target, 0, high)
    if nums[p] == target
        return true
    if target < nums[0]
        low = p
    else
        low = 1
        if p != 0
            high = mid -1
    for low <= high
        mid = (low+high) /2
        if nums[mid] == target
            return true
        elseif nums[mid] < target
            low = mid +1
        else
            high = mid -1
    return false
    
FIND-ROTATE-POSITION(nums, target, low, high)
    for low <= high
        mid = (low+high) /2
        if nums[mid] == target
            return mid
        if nums[0] < nums[mid] && mid +1 < nums.length
            if nums[0] >= nums[mid+1]
                return mid +1
            else
                low = mid +1
        elseif nums[0] > nums[mid]
            high = mid -1
        elseif nums[0] == nums[mid]
            r = FIND-ROTATE-POSITION(nums, target, low, mid -1)
            if r > 0
                return r
            return FIND-ROTATE-POSITION(nums, target, mid +1, high)
        else
            return 0
    return 0
```