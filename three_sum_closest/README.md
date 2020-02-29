### 最接近的三数之和（3Sum Closest）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/3sum-closest/)

#### 题目
给定一个包括 n 个整数的数组`nums`和一个目标值`target`。找出`nums`中的三个整数，使得它们的和与`target`最接近。返回这三个数的和。假定每组输入只存在唯一答案。
```
例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.

与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).
```

#### 思路
这题与求三数和为 0 的题解法几乎一致，大致步骤如下：

- 如果`nums[i] == nums[i-1]`，那么`nums[i]`存在解必然与`nums[i-1]`重复，要么不存在解，则可以直接跳过；
- 如果`nums[i] > target`，则可以直接跳出循环，因为后面的数都比`nums[i]`大，不可能再遇到三数和更接近`target`的情况；
- 如果`sum = target`，则直接返回`target`；
- 如果`sum-target`，只需要向前移动右指针；如果`sum < target`，只需要向后移动左指针；
- 如果`nums[i+1] == nums[i]`或`nums[r-1] == nums[r]`会导致重复计算，则可以直接跳过。

#### 伪代码
THREE-SUM-CLOSEST(nums, target)
    sort(nums)
    res = target > 0 ? math.MinInt32 : math.MaxInt32
    for i = 0 to nums.length -3
        if nums[i] > target
            break
        l, r = i +1, nums.length -1
        while l < r
            sum = nums[i] + nums[l] + nums[r]
            if sum == target
                return target
            if math.Abs(sum - target) < math.Abs(res - target)
                res = sum
                if sum < target
                    while l < r && nums[l+1] == nums[l]
                        l ++
                    l ++
                else
                    while l < r && nums[r-1] == nums[r]
                        r --
                    r --
    return tes
                