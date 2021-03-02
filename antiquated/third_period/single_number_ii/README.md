### 只出现一次的数字 II（Single Number II）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/single-number-ii/)

#### 题目
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现了三次。找出那个只出现了一次的元素。

说明：  
你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

示例 1：
```
输入: [2,2,3,2]
输出: 3
```
示例 2：
```
输入: [0,1,0,1,0,1,99]
输出: 99
```

#### 思路
线性时间复杂度很好做，用一个 hashMap 就能解决，但是达不到不使用额外空间的要求。

#### 思路二（位运算）
[位运算题解](https://leetcode-cn.com/problems/single-number-ii/solution/single-number-ii-mo-ni-san-jin-zhi-fa-by-jin407891/)

#### 伪代码
```
SINGLE-NUMBER(nums []int)
    m = new Map()
    for i = 0 to nums.length -1
        m[nums[i]] ++
    keys = m.keys()
    for i = 0 to keys.length -1
        if m[keys[i]] == 1
            return keys[i]
```