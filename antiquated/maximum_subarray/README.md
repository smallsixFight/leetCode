### 最大子序和（Maximum Subarray）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/maximum-subarray/)

#### 题目
给定一个整数数组`nums`，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例：
```
输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
```
进阶：  
如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。

#### 思路
这题使用动态规划来实现，使用自底向上方式来实现，需要找到动态规划的那个 k，通过分析后发现 k 的特性有：

- 最大和子数组一定包含在`nums[0:k]`或`nums[k+1:nums.length -1]`中。

另外，还能发现最大子数组具备下列特性：

- 最大和子数组长度大于 1 时，第一个一定不为负数；
- 当遇到使得直至目前的最大和值小于 0 的数字时，那么目前的子数组的最大和已经达到最大值，该值的索引即为 k；
- 应从 k+1 开始寻找下一个 k‘ 和子数组最大和；

#### 涨姿势
`maxSubArray2()`是在 LeetCode 学到的写法，我对动态规划的应用还是不够啊。

#### 伪代码
```
MAX-SUBARRAY(nums)
    if nums.length == 0
        return 0
    p1, p2 = 0, 1
    maxSum = nums[0]
    temp = maxSum
    for p2 to nums.length -1
        if nums[p1] < 0 && nums[p2] > nums[p1]
            p1 = p2
            temp = nums[p1]
        else
            temp += nums[p2]
        if temp > maxSum
            maxSum = temp
        p2 ++
        if temp < 0     // tempSum 小于 0 处理
            p1 = p2
            temp = 0
    return maxSum
```