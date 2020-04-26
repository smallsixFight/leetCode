###  乘积最大子数组（Maximum Product Subarray）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/maximum-product-subarray/)

#### 题目
给你一个整数数组`nums`，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字）。

示例 1：
```
输入: [2,3,-2,4]
输出: 6
解释: 子数组 [2,3] 有最大乘积 6。
```
示例 2：
```
输入: [-2,0,-1]
输出: 0
解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
```

#### 思路一
这道题可以用 DP 来做，可以发现，第 n 个元素的能够获取的最大积取决与它所有的积跟前 n-1 个的最大积的比较，而第 n 个数有 n-1 个积。

提交后，最后一个测试案例观望直接内存溢出给我炸了。

#### 思路二
这个思路是在看完 [画解算法的题解](https://leetcode-cn.com/problems/maximum-product-subarray/solution/hua-jie-suan-fa-152-cheng-ji-zui-da-zi-xu-lie-by-g/) 后个人的剖析过程记录。

先把题目做一个减法，假设数组的元素都是非负数，那么只需要获取第 n 位能够获取的最大值，而第 n 位的最大值有公式：`curMax(n) = curMax(n-1) * nums[n]`，如果当前位置的最大值大于前 n-1 位的最大值，则替换。

示例：
```
arr:    [1, 2, 3, 0, 1, 2, 3, 4]
curMax:  1, 2, 6, 0, 1, 2, 6, 24
max:     1, 2, 6, 6, 6, 6, 6, 24
```

接着，把负数考虑进来，那么最大值就会变成最小值，但是反向一下，第 n-1 位的最小值与第 n 位的值相乘就能得到当前位置的最大值，那么稍微变通一下逻辑，如果遇到负数，第 n-1 位的最小值和最大值交换一下即可，示例如下：
```
arr:    [1, 2, -2, -3,  0, -1,  2,  3]
curMax:  1, 2, -2, 12,  0,  0,  2,  6
curMin:  1, 1, -4, -3,  0, -1, -2, -6
max:     1, 2,  2, 12, 12, 12, 12, 12
```

#### 伪代码
``````
MAX-PRODUCT-TWO(nums)
    if nums.length == 0
        return 0
    maxVal, preMax, preMin = nums[0], 1, 1
    for i = 0 to nums.length -1
        if nums[i] < 0
            preMax, preMin = preMin, preMax
        preMax = max(preMax * nums[i], nums[i])
        preMin = min(preMin * nums[i], nums[i])
        maxVal = max(preMax, maxVal)
    return maxVal

MAX-PRODUCT(nums)
    if nums.length == 0
        return 0
    dp = new Array()
    res = nums[0]
    for i = 0 to nums.length -1
        dp[i][i] = nums[i]
        if num[i] > res
            res = nums[i]
    for i = 0 to nums.length -1
        for j = i +1 to nums.length -1
            dp[i][j] = dp[i][j-1] * dp[j][j]
            if dp[i][j] > res
                res = dp[i][j]
    return res