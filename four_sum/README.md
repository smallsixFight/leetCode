### 四数之和（4Sum）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/4sum/)

#### 题目
给定一个包含 _n_ 个整数的数组`nums`和一个目标值`target`，判断`nums`中是否存在四个元素 _a，b，c_ 和 _d_，使得 _a + b + c + d_ 的值与`target`相等，找出所有满足条件且不重复的四元组。

**注意：**  
答案中不可以包含重复的四元组。

**示例：**  
```
给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。

满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
```

#### 思路
这道题可以按照三数和那道题的解法来做，只需要前两个数看作一个整体，那就跟三数和很相似了。

由于 target 可以为正负数，跳过的条件需要进行调整。

#### 伪代码
```
FOUR-SUM(nums, target)
    if nums.length < 4
        return empty array
    sort(nums)
    res = new Array()
    for i = 0 to nums.length -4
        if i > 0 && nums[i] == nums[i-1]
            continue
        for j = i +1 to nums.length -3
            l, r = j +1, nums.length -1
            if nums[i] + nums[j] > target && nums[j] > 0
                break
            if j > 1 && nums[j] == nums[j-1]
                continue 
            while l < r
                sum = nums[i] + nums[j] + nums[l] + nums[r]
                if sum == target
                    res.push([nums[i], nums[j], nums[l], nums[r]])
                    for l < r && nums[l+1] == nums[l]
                        l ++
                    for l < r && nums[r-1] == nums[r]
                        r --
                    l ++
                    r --
                elseif sum > target
                    r --
                else
                    l ++
    return res
```