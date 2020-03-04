### 寻找两个有序数组的中位数（Median of Two Sorted Arrays）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/median-of-two-sorted-arrays/)

#### 题目
给定两个大小为 m 和 n 的有序数组`nums1`和`nums2`。

请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设`nums1`和`nums2`不会同时为空。

示例 1：
```
nums1 = [1, 3]
nums2 = [2]

则中位数是 2.0
```

示例 2：
```
nums1 = [1, 2]
nums2 = [3, 4]

则中位数是 (2 + 3)/2 = 2.5
```

#### 思路
如果在一个有序数组中，那么它的中位数一定是`nums[mid+1]`或`(nums[mid]+nums[mid+1])/2`。如果是两个两序数组，第一个想到办法就是将两个数组合并成一个数组，再去获取中位数，然而这样很明显不符合题目的时间复杂度要求。

那么可以一边排序一边查找`mid`，由于永远都只需要遍历`(nums1.length +nums2.length)/2`次，比前面减少了一半，不过时间复杂度依旧是 O(m+n)。

下面是在看了官方题解后自己整理的思路。

先假设数组`A`和`B`都不为空，分别把数组`A`和`B`从中间拆分成`A[0:i-1]、A[i:]、B[0:j-1]、B[j:]`，然后将`A[0:i-1]`和`B[0:j-1]`组成 part_1，将`A[i:]`和`B[j:]`组成 part_2。要获取中位数，则需要满足下面的条件：

- MAX(part_1) < MIN(part_2)；
- 当((m+n) %2 == 0)，`len(A[0:i-1])+len(B[0:j-1]) == len(A[i:])+len(B[j:])`；
- 当((m+n) %2 == 1)，则满足`len(A[0:i-1])+len(B[0:j-1])+1 == len(A[i:])+len(B[j:])`或`len(A[0:i-1])+len(B[0:j-1])` == `len(A[i:])+len(B[j:])+1`。

下一步就是如何获取 i 和 j 了。为了满足题目对数组时间复杂度的要求，那么想到就是二分查找，那么可以通过 i 来得到 j 吗？这是可以的。

因为中位数的左右两边长度相等，那么 i +j == (i-1) + (j-1)；那么`i+j == (m+n)/2`或`(i+j) == (m+n+1)/2`，就可以得到`j = (m+n+1)/2 -i`或`j = (m+n)/2 -i`。

结合前面，我们只需要对`i`或`j`进行二分查找。

i、j 都不能小于 0，当对`i`进行二分查找的时候，只需要考虑不能让`j<0`的情况出现，因为`i`最大值为`m`，那么`m<n`一定要作为前提条件成立，所以数组`A`必为长度较小的数组。

临界值描述：

- 由于`m<=n`为前提，那么`j`不可能大于`n`，且`j==n`仅在`m == n and i == 0`的情况下成立；
- 由于`m<=n`为前提，那么`j`不可能小于 0，且`j==n`仅在`m == n and i == m`的情况下成立；
- 当`i == 0`时，中位数为`B[j-1]`或`(B[j-1]+B[j])/2`；
- 当`i == m and m < n`时，中位数为`B[j-1]`或`(B[j-1]+B[j])/2`；
- 当`i == m and m == n`时，中位数为`(A[i]+B[j])/2`。

那么，可以尝试写下伪代码了。

#### 伪代码
```
FIND-MEDIAN-SORTED-ARRAYS(nums1, nums2)
    m, n = nums1.length, nums2.length
    if m > n
        m, n = n, m
        nums1, nums2 = nums2, nums1
    iMin, iMax, halfLen = 0, m, (m+n+1)/2
    while iMin <= iMax
        i = (iMin+iMax)/2
        j = halfLen -i
        if i < iMax && nums2[j-1] > nums1[i]
            iMin = i +1
        elseif i > iMin && nums1[i-1] > nums2[j]
            iMax = i -1
        else
            max1 = 0
            if i == 0
                max1 = nums2[j-1]
            elseif j == 0   // m == n
                max1 = nums1[i-1]
            else
                if nums1[i-1] > nums2[j-1]
                    max1 = nums1[i-1]
                else
                    max1 = nums2[j-1]
                if (m+n) %2 == 1
                    return max1
                max2 = 0
                if i == m
                    max2 = nums2[j]
                elseif j == n
                     max2 = nums1[i]
                else
                    if nums1[i] < nums2[j]
                        max2 = nums1[i]
                    else
                        max2 = nums2[j]
                return (max1 + max2)/2
    return 0
```

#### 题外话
按照伪代码来写的代码提交后运行时间为 24ms，但在自定义`max()`和`min()`函数后，直接提升到 8ms、16ms、20ms 等，也许这就是优化吧。