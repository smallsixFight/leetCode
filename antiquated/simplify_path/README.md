### 简化路径（Simplify Path）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/simplify-path/)

#### 题目
以 Unix 风格给出一个文件的**绝对路径**，你需要简化它。或者换句话说，将其转换为规范路径。

在 Unix 风格的文件系统中，一个点（.）表示当前目录本身；此外，两个点 （..） 表示将目录切换到上一级（指向父目录）；两者都可以是复杂相对路径的组成部分。更多信息请参阅：Linux / Unix中的绝对路径 vs 相对路径

请注意，返回的规范路径必须始终以斜杠`/`开头，并且两个目录名之间必须只有一个斜杠`/`。最后一个目录名（如果存在）不能以`/`结尾。此外，规范路径必须是表示绝对路径的**最短**字符串。

示例 1：
```
输入："/home/"
输出："/home"
解释：注意，最后一个目录名后面没有斜杠。
```
示例 2：
```
输入："/../"
输出："/"
解释：从根目录向上一级是不可行的，因为根是你可以到达的最高级
```
示例 3：
```
输入："/home//foo/"
输出："/home/foo"
解释：在规范路径中，多个连续斜杠需要用一个斜杠替换。
```
示例 4：
```
输入："/a/./b/../../c/"
输出："/c"
```
示例 5：
```
输入："/a/../../b/../c//.//"
输出："/c"
```
示例 6：
```
输入："/a//b////c/d//././/.."
输出："/a/b/c"
```

#### 思路
这道题可以使用链表来保存有效的结果集，最后将结果集转换为字符串。

以下是对不同情况的处理方式整理：

- path 首个字符不是`/`时需要先入一个`/`；
- 如果当前字符为`/`且下一个字符为`/`，那么忽略该字符；
- 如果当前字符为`.`且下一个字符为`/`，则忽略这两个字符；
- 如果当前字符为`.`且下两个字符为`.`和`/`，则将移除链表最后的目录，如果链表最后的目录为根目录，则忽略；然后跳过这两个字符；
- 其他情况都加入到列表中。

#### 题外话
自己对目录命名的不熟悉加上被题目带了一下，导致以为遇到`.`就是当前目录可以忽略，遇到`..`就返回上层目录，其实应该是遇到`./`和`../`才对。

#### 伪代码
```
SIMPLIFY-PATH(path)
    list = new ArrayList()
    for i = 0 to path.length -1
        while path[i] == '/' && i +1 < path.length && path[i+1] == '/'
            i ++
        p = []byte
        if path[i] != '/'
            p.push('/')
        else
            p.push(path[i])
            i ++
        pointNum = 0
        while i < path.length
            if path[i] == '.'
                pointNum ++
            else if path[i] == '/'
                break
            p.push(path[i])
            i ++
        if pointNum == 2 && p.length == 3       // 两个点（加上斜杠
            if list.length > 0 && list[list.length-1] != "/'
                list = list[:list.length-1]
        else if pointNum == 1
            if p.length > 2
                list.push(p.toString())
        else if p.length > 1
            list.push(p.toString())
    s = list.toString()
    if s == ""
        return "/"
    return s
```