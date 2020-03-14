### 下一个排列（Next Permutation）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/next-permutation/)

#### 题目
实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。

如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。

必须`原地`修改，只允许使用额外常数空间。

以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
```
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1
```

#### 思路
题目给出的例子其实不是很明确，所以我用一个较长的例子来说明，`[1, 9, 8, 4, 7, 6, 3, 2]`，这个数组执行后得到的结果应该是`[1, 9, 8, 6, 2, 3, 4, 7]`。
分步来说明我的思路：  
- 从后向前遍历，找出第一个比当前值小的值的索引，分别记录为`p1`和`p2`，进行下一次遍历比较；
- 当后续的遍历比较也存在比当前值小的值，而索引位置小于等于`p1`时，忽略它，进行下一次遍历比较；否则，更新`p1`和`p2`的值。
- 遍历结束后，交换`p1`和`p2`上的值，然后对`p1`后的元素做升序排序；
- 如果遍历结束后不需要交换任何的值，那么需要对整个数组做升序排序；
- 结合上面的情况，`p1`的初始值可以设为 -1，当不为 -1 时，进行值的交换，然后对`nums[p1+1:]`的元素做升序排序。

由于题目要求，排序我用最大堆排序来实现。

这道题不同考虑数组第一位为 0 需要忽略的情况。

#### 题外话
为什么要用堆排序，其实就两个原因：  
1. 堆排序的时间复杂度小，且空间复杂度为 O(1)；
2. 复习一下堆排序。

#### 伪代码
```
NEXT-PERMUTATION(nums)
    if nums.length == 0
        return
    p1, p2 = -1, 0
    for i = nums.length -1 to 1
        for j = i -1 to 0
            if nums[i] > nums[j] && (j > p1 || p1 == -1)
                p1 = j
                p2 = i
                break
    if p1 != -1
        temp = nums[p1]
        nums[p1] = nums[p2]
        nums[p2] = nums [p1]
    MAX-HEAP-SORT(nums[p1+1:])

MAX-HEAP-SORT(nums)
    for i = nums.length /2 to 0         // 建堆
        MAX-HEAPIFY(nums, i)
    for lastIdx = nums.length -1 to 0    // 缩减堆数量并维护堆性质
        nums[0], nums[lastIdx] = nums[lastIdx], nums[0]
        MAX-HEAPIFY(nums[0:lastIdx], 0)

MAX-HEAPIFY(arr, i)
    l, r = 2*i +1, 2*i +2
    largest = 0
    if l < arr.length && arr[l] > arr[i]
        largest = l
    else
        largest = i
    if r < arr.length && arr[r] > arr[largest]
        largest = r
    if largest != i
        arr[i], arr[largest] = arr[largest], arr[i]
        MAX-HEAPIFY(arr, largest)
```