### 三数之和（3Sum）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/3sum/)

#### 题目
给定一个包含 n 个整数的数组`nums`，判断`nums`中是否存在三个元素 a，b，c ，使得 a + b + c = 0。找出所有满足条件且不重复的三元组。

**注意：**
答案中不可以包含重复的三元组。

示例：
```
给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
```
#### 题外话
我太菜了，这道题没做出来，去重的问题真的考到我了。LeetCode 这题最高阅读量的居然还是个垃圾题解，前面写一堆都不是正确答案就敢放，下一个标题居然是“再优化”，好歹答案正确才能提优化吧（我在电脑前骂骂咧咧

思路记录的是看了排序+左右双指针的题解后大概整理的记录，排序后去重就变得很简单。

#### 思路
先把数组按非降序排序，从头开始遍历，使用双指针来计算三数和，有序后可以有下列的一些规则：

- 如果`nums[i]` > 0，则可以直接跳出循环，因为后面的数都比`nums[i]`大，不可能再遇到三数和为零的情况；
- 如果`nums[i] == nums[i-1]`，若`nums[i-1]`存在解，那么`nums[i]`会导致结果重复，若`nums[i-1]`不存在解，`nums[i]`肯定也不存在解，则可以直接跳过；
- 如果 sum = 0，若`nums[l+1] == nums[l]`或`nums[r-1] == nums[r]`都会导致结果重复，可以直接跳过。 
- 如果 sum > 0，只需要向前移动右指针；如果 sum < 0，只需要向后移动左指针。

#### 伪代码
```
THREE-SUM(nums)
    if nums.length < 3
        return empty array
    sort(nums)
    res = new array()
    for i = 0 to nums.length - 3
        if nums[i] > 0
            break
        if i > 0 && nums[i] == nums[i-1]
            continue
        l, r = i +1, nums.length -1
        while l < r
            sum = nums[i] + nums[l] + nums[r]
            if sum == 0
                res.push(nums[i], nums[l], nums[r]])
                for l < r && nums[l+1] == nums[l]
                    l ++
                for r > l && nums[r-1] == nums[r]
                    r --
                l ++
                r --
            elseif sum > 0
                r --
            else
                l ++
        return res
```