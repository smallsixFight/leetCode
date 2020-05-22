### 跳跃游戏 II（Jump Game II）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/jump-game-ii/)

#### 题目
给定一个非负整数数组，你最初位于数组的第一个位置。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

你的目标是使用最少的跳跃次数到达数组的最后一个位置。

示例：
```
输入: [2,3,1,1,4]
输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。
     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
```
说明：

假设你总是可以到达数组的最后一个位置。

#### 思路
这道题与之前做的第一道不同，之前那道是判断能否到达最后一个位置，而这道题直接给出**总是可以到达数组的最后一个位置**，也就是数组不存在 0，不过鉴于 LeetCode 的题目有时候描述确实坑（也许是我个人的理解能力有限），个人还是考虑进去了。

这道题的解题思路基本上可以完全照搬前面那道题的思路。

只需要遍历当前获取当前位置能够跳到的所有位置中能够跳到最远的位置即可。

#### 伪代码
```
JUMP(nums)
    if nums.length < 2
        return 0
    res = 0
    curIdx := 0
    for nums[curIdx] > 0
        if curIdx + nums[curIdx] >= nums.length -1
            return res
        res ++
        p1 = curIdx
        furthest = 0
        for i = nums[p1] to 1
            if p1 + i + nums[p1+i] > furthest
                furthest = p1 + i + nums[p1+i]
                curIdx = p1 +i
    return -1
```