### 盛最多水的容器（Container With Most Water）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/container-with-most-water/)

#### 题目
给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

**说明：**
你不能倾斜容器，且 n 的值至少为 2。  
![question_11.jpg](https://i.loli.net/2020/02/29/7OVyFzHpQBa1U9N.jpg)

图中垂直线代表输入数组`[1,8,6,2,5,4,8,3,7]`。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。

示例：
```
输入: [1,8,6,2,5,4,8,3,7]
输出: 49
```

#### 思路
额，先翻译一下就是：求 x 轴和 y 轴能达到的最大面积。

这题可以使用暴力法求出所有的值两两之间的面积，比较获取最大值。

另外，可以发现，面积的最大值收到两条边那条长度较小的影响，所以将宽度缩小，读取另外一条边来求面积时，可能得到更大的面积。
那么，可以使用双指针法来进行移动。

#### 伪代码
```
MAX-AREA(height)
    maxArea, p1, p2 = 0, 0, height.length -1
    for p1 < p2
        h = 0
        if height[p1] > height[p2]
            h = height[p2]
            p2 --
        else
            h = height[p1]
            p1 ++
        area = (p2 - p1 +1) * h
        if area > maxArea
            maxArea = area
    return maxArea

MAX-AREA2(height)
    p1, p2 = 0, 1
    maxArea = 0
    for p1 to height.length -2
        for p2 to height.length -1
            if height[p1] > height[p2]
                h = height[p2]
            else
                h = height[p1]
            area = (p2 - p1) * h
            if area > maxArea
                maxArea = area
    return maxArea
```