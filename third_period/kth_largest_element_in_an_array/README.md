### 数组中的第 K 个最大元素（Kth Largest Element in an Array）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/kth-largest-element-in-an-array/)

#### 题目
在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

示例 1：
```
输入: [3,2,1,5,6,4] 和 k = 2
输出: 5
```
示例 2：
```
输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
输出: 4
```

说明：  
你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。

#### 思路一
这道题就是排序，然后返回第 k 大的值。调用类库两行代码就做完了。

这道我写了一个归并排序来完成。

#### 思路二
由于归并排序的空间复杂度是 O(n)，所以又练习写了一下堆排序。

堆排序主要要先理解最大堆和最小堆两个概念。

#### 伪代码
```
FIND-KTH-LARGEST(nums, k)
    nums = MERGE-SORT(nums)
    return nums[nums.length - k]

MERGE-SORT(nums)
    if nums.length < 1
        return nums
    a1 = MERGE-SORT(nums[:nums.length/2])
    a2 = MERGE-SORT(nums[nums.length/2:])
    res = new Array()
    i, j = 0, 0
    idx = 0
    while i < a1.length || j < a2.length
        if i == a1.length
            res[idx] = a2[j]
            j ++
        elseif j == a2.length
            res[idx] = a1[i]
            i ++
        elseif a1[i] < a2[j]
            res[idx] = a1[i]
            i ++
        else
            res[idx] = a2[j]
            j ++
        idx ++
    return res
```