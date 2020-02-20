### 两数之和（Two Sum）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/two-sum/)

#### 题目
给定一个整数数组 nums 和一个目标值`target`，请你在该数组中找出和为目标值的那`两个`整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

#### 思路
第一种是暴力破解，使用两个 for 循环（一个内嵌），依次相加匹配。

第二种是用 Map 这个数据结构，key 放整数，value 放索引。先把数组全部转换为 Map，然后再查找是否有符合的 key。

#### 伪代码
```
FirstMethod(array, target) 
    for i = 0 to array.length -2
        for j = i +1 to array.length -1
            if array[i] == target - array[j]
                return [i, j]
    return NIL
```
```
SecondMethod(array, target)
    m = new Map()
    for i = 0 to array.length -1
        if m[target - array[i]] != NIL
            return [m[target - array[i]], i]
        m[array[i]] = i     // 避免重复使用
    return NIL
```
