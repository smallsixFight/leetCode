### 移除元素（Remove Element）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/remove-element/)

#### 题目
给定一个数组 nums 和一个值 val，你需要`原地`移除所有数值等于 val 的元素，返回移除后数组的新长度。

不要使用额外的数组空间，你必须在`原地修改输入数组`并在使用 O(1) 额外空间的条件下完成。

元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

示例 1:
```
给定 nums = [3,2,2,3], val = 3,

函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。

你不需要考虑数组中超出新长度后面的元素。
```
示例 2:
```
给定 nums = [0,1,2,2,3,0,4,2], val = 2,

函数应该返回新的长度 5, 并且 nums 中的前五个元素为 0, 1, 3, 0, 4。

注意这五个元素可为任意顺序。

你不需要考虑数组中超出新长度后面的元素。
```

#### 思路
这道题可以使用双指针法来解决。指针 p1、p2 初始均指向`nums[0]`。  
当`nums[p1] == val && nums[p2] != val`时 交换`nums[p1]`和`nums[p2]`的值，然后 p1、p2 向后移动，长度加一。  
当`nums[p1] == val && nums[p2] == val`时 p2 向后移动。
当`nums[p1] != val`时，p1 向后移动，长度加一；如果`nums[p2] == val || p2 == p1 +1`，p2 向后移动。
当 p2 遍历结束时，直接返回。

#### 伪代码
```
REMOVE-ELEMENT(nums, val)
    p1, p2 = 0, 0
    l = 0
    while p2 < nums.length
        if nums[p1] == val
            if nums[p2] != val
                exchange(nums[p1], nums[p2])
                l += 1
                p1 += 1
            p2 += 1
        else
            l += 1
            p1 += 1
            if nums[p2] == val || p1 >= p2
                p2 += 1
    return l
```