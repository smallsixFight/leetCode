### 缺失的第一个正数（First Missing Positive）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/first-missing-positive/)

#### 题目
给你一个未排序的整数数组，请你找出其中没有出现的最小的正整数。

示例 1：
```
输入: [1,2,0]
输出: 3
```
示例 2：
```
输入: [3,4,-1,1]
输出: 2
```
示例 3：
```
输入: [7,8,9,11,12]
输出: 1
```

提示：  
你的算法的时间复杂度应为`O(n)`，并且只能使用常数级别的额外空间。

#### 思路
`O(n)`的时间复杂度已经把用排序处理堵掉了，常数级别的额外空间要求则不能使用哈希表来处理。

不过可以利用哈希表的思维进一步思考，可以将原数组的作为哈希表来使用，大致思路如下：

- 先遍历一次数组，如果值为 1，则计数器加一，跳出循环；
- 如果计数器为 0，则直接返回 1；如果计数器不为 0，但是数组长度为 1，则直接返回 2；
- 遍历整个数组，将所有大于数组长度和小于 1 的值全部转为 1；
- 遍历整个数组，如果值为正数，则将数组对应索引的值转为负数；
- 从第二位遍历整个数组，如果值为正数，则返回该值；
- 遍历结束后，如果数组第一位为正数，返回该值；否则返回数组长度 +1。

#### 题外话
总感觉困难题的难度在于需要更加深入的思考，现在还是有种我在第二层，题解在第五层的感觉。

#### 伪代码
```
FIRST-MISSING-POSITIVE(nums)
    var hasOne = false
    for i 0 to nums.length -1
        if nums[i] == 1
            hasOne = true
            break
    if !hasOne
        return 1
    if nums.length == 1
        return 2
    for i = 0 to nums.length -1
        if nums[i] < 1 || nums[i] > nums.length
            nums[i] = 1
    for i = 0 to nums.length -1
        if nums[i] < nums.length
            nums[nums[i]] = - math.Abs(nums[nums[i]])
        else
            nums[0] = -math.Abs(nums[0])
    for i = 1 to nums.length -1
        if nums[i] > 0
            return nums[i]
    if nums[0] > 0
        return nums[0]
    return nums.length +1
```