###  寻找峰值（Find Peak Element）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/find-peak-element/)

#### 题目
峰值元素是指其值大于左右相邻值的元素。

给定一个输入数组`nums`，其中`nums[i] ≠ nums[i+1]`，找到峰值元素并返回其索引。

数组可能包含多个峰值，在这种情况下，返回任何一个峰值所在位置即可。

你可以假设`nums[-1] = nums[n] = -∞`。

示例 1：
```
输入: nums = [1,2,3,1]
输出: 2
解释: 3 是峰值元素，你的函数应该返回其索引 2。
```
示例 2：
```
输入: nums = [1,2,1,3,5,6,4]
输出: 1 或 5 
解释: 你的函数可以返回索引 1，其峰值元素为 2；
     或者返回索引 5， 其峰值元素为 6。
```

说明：

你的解法应该是 O(logN) 时间复杂度的。

#### 思路
时间复杂度的要求就是要用二分法。再来看看题目。

由于`nums[-1] = nums[n] = -∞`并且`nums[i] = nums[i+1]`，那么数组必然至少存在一个峰值。这样就可以有如下使用二分法的判断条件：

- 如果`nums[i-1] < nums[i] < nums[i+1]`，那么峰值在 i 的右侧；
- 如果`nums[i-1] > nums[i] > nums[i+1]`，那么峰值在 i 的左侧；
- 如果`nums[i-1] > nums[i] && nums[i] < nums[i+1]`，那么 i 两侧都有峰值，随便挑一边即可。

#### 伪代码
```
FIND-PEAK-ELEMENT(nums)
    if nums.length < 2
        return 0
    low, high = 0, nums.length -1
    for low <= high
        mid = (low + high) >> 1
        if (mid == 0 && nums[0] > nums[1]) || (mid == nums.length -1 && nums[nums.length-1] > nums[nums.length-2])
            return mid
        if mid > 0 && mid < nums.length && nums[mid-1] < nums[mid] && nums[mid] < nums[mid+1]
            return mid
        if mid == 0 || (nums[mid-1] < nums[mid] && nums[mid] < nums[mid+1])
            low = mdi +1 
        else
            high = mid -1
    return 0
```