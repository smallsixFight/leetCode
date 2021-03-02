### 接雨水（Trapping Rain Water）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/trapping-rain-water/)

#### 题目
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

![rainwatertrap.png](https://i.loli.net/2020/05/14/eISrWRZYJuxc8jX.png)

上面是由数组`[0,1,0,2,1,0,1,3,2,1,2,1]`表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢 Marcos 贡献此图。

示例：
```
输入: [0,1,0,2,1,0,1,3,2,1,2,1]
输出: 6
``` 

#### 思路
使用一个数组记录当前位置与最近第二高的高度差距，使用两个指针分别记录开始位置和当前遍历到的位置。

当遇到高度大于等于开始位置的柱子的高度时，计算当前可以存的雨水数值，然后移动开始指针到当前遍历到的位置，设置为新的起点。

第一轮遍历结束后，从尾部开始，遍历找到第一个开始凹进去的起点，如果该起点不是前面起点的位置，则遍历获取两者之间的存储的水。

#### 伪代码
```
TRAP(height)
    if height.length < 3
        return 0
    p1, p2 := 0, 0
    res = 0
    while p1 + 1 < height.length    // 找到起点
        if height[p1+1] < height[p1]
            p2 = p1 +1
            break
        p1 ++
    while p2 < height.length
        while p2 < height.length -1 && height[p2] < height[p1]
            p2 ++
        if height[p2] >= height[p1]
            for i = p1 +1 to p2 -1
                res += height[p1] - height[i]
            p1 = p2
        p2 ++
    if p1 != p2 // 如果不是结束的位置，则需要从尾部遍历一次
        while p2 > p1 && height[p2] <= height[p2-1]
            p2 --
        for i = p2 -1 to p1 +1
            if height[i] < height[p2]
                res += height[p2] - height[i]
            else
                p2 = i
    return res
```