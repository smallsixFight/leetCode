### 打家劫舍 II（House Robber II）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/house-robber-ii/)

#### 题目
你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都围成一圈，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。

示例 1：
```
输入: [2,3,2]
输出: 3
解释: 你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。
```
示例 2：
```
输入: [1,2,3,1]
输出: 4
解释: 你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。
     偷窃到的最高金额 = 1 + 3 = 4 。
```

#### 思路
这道题明显是一道 DP 题，这题的重点在于最后一家偷不偷，所以可以分为从第一家开始偷和从第二家开始偷两种情况，然后返回它们的最大值。

#### 伪代码
```
ROB(nums []int)
    if nums.length == 0
        return 0
    elseif nums.length == 1
        return nums[0]
    if nums.length == 2
        return MAX(nums[0], nums[1])
    return max(TRAVERSE(0, nums.length-2, nums), TRAVERSE(1, nums.length-1, nums))
    
TRAVERSE(a, b, nums)
    dp = [b-a+1]int{}
    dp[0] = nums[a]
    dp[1] = max[nums[a], nums[a+1]]
    for i = 2 to b-a
        dp[i] = MAX(nums[a+1] + dp[i-2], dp[i-1])
    return dp[b-a]
```