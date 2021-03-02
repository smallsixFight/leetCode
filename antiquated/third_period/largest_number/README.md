### 第 N 高的薪水（Nth Highest Salary）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/largest-number/)

#### 题目
给定一组非负整数，重新排列它们的顺序使之组成一个最大的整数。

示例 1：
```
输入: [10,2]
输出: 210
```
示例 2：
```
输入: [3,30,34,5,9]
输出: 9534330
```
说明：输出结果可能非常大，所以你需要返回一个字符串而不是整数。

#### 思路
这道题的思路其实挺清晰的，就是实现的问题的而已。解题思路大致如下：

- 把所有的数字当作字符串看，从头开始比较第一个数字，大的排在前面；
- 然后依次比较相同前缀的后一个数字大小，大的排在前面，以此类推；
- 相同前缀的长的剩余部分要跟短的递归比较，要一直比短的大；
- 就像 SQL 语句的按多个字段排序那样。

由于个人做 LeetCode 的原则是尽量不要用标准库，所以代码就比较长了，不过应该还是比较容易看懂的。

#### 伪代码
```
LARGEST-NUMBER(nums)
    arr = new Array()
    for i = 0 to nums.length -1
        if nums[i] == 0
            arr.add("0")
            continue
        s = new Stack()
        for nums[i] > 0
            s.push(nums[i]%10 + '0')
            nums[i] /= 10
        a = ""
        for !s.isEmpty()
            a += s.pop()
        arr.push(a)
    CUSTOM-SORT(arr)
    res = ""
    for i = 0 to arr.length -1
        res += arr[i]
    return res
    
CUSTOM-SORT(arr)
    if arr.length < 2
        return arr
    arr1, arr2 = arr[:arr.length/2], arr[arr.length/2:]
    arr1 = CUSTOM-SORT(arr1)
    arr2 = CUSTOM-SORT(arr2)
    res = new Array()
    i, j, idx = 0, 0, 0
    while i < arr1.length || j < arr2.length
        if i == arr1.length
            res.add(arr2[j])
            j ++
        elseif j == arr2.length
            res.add(arr1[i])
            i ++
        elseif arr1[i] == arr2[j]
            res.add(arr1[i])
            res.add(arr2[j])
            i ++
            j ++
        elseif arr1[i][0] > arr2[j][0]
            res.add(arr1[i])
            i ++
        elseif arr[i][0] < arr2[j][0]
            res.add(arr2[j])
            j ++
        else
            if IS-BIGGER(arr1[i], arr2[j])
                res.add[arr1[i]]
                i ++
            else
                res.add(arr2[j])
                j ++
                    
IS-BIGGER(s1, s2)
    is s1.length == 0 || s2.length == 0
        return true
    i = 0
    while true
        if i == s1.length
            return IS-BIGGER(s1, s2[i:])
        elseif i == s2.length
            return IS-BIGGER(s1[i:], s2)
        if s1[i] > s2[i]
            return true
        else s1[i] < s2[i]
            return false
        i ++
```