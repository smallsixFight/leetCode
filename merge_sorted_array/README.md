### 合并两个有序数组（Merge Sorted Array）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/merge-sorted-array/)

#### 题目
给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。

*说明：*
- 初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
- 你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。

示例：
```
输入:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

输出: [1,2,2,3,5,6]
```

#### 思路
本题使用双指针法从后向前遍历比较，较大的值向后移动。

#### 伪代码
MERGE(nums1, m, nums2, n)
    for m + n > 0
        if m > 0 && n > 0
            if nums1[m-1] > nums2[n-1]
                nums1[m+n-1] = nums1[m-1]
                m --
            else
                nums1[m+n-1] = nums1[n-1]
                n --
        else if n > 0
            nums1[n-1] = nums2[n-1]
            n --
        else
            break