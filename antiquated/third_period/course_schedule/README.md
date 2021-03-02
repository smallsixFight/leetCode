### 课程表（Course Schedule）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/course-schedule/)

#### 题目
你这个学期必须选修`numCourse`门课程，记为`0`到`numCourse-1`。

在选修某些课程之前需要一些先修课程。 例如，想要学习课程`0`，你需要先完成课程`1`，我们用一个匹配来表示他们：`[0,1]`

给定课程总量以及它们的先决条件，请你判断是否可能完成所有课程的学习？

示例 1：
```
输入: 2, [[1,0]] 
输出: true
解释: 总共有 2 门课程。学习课程 1 之前，你需要完成课程 0。所以这是可能的。
```
示例 2：
```
输入: 2, [[1,0],[0,1]]
输出: false
解释: 总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0；并且学习课程 0 之前，你还应先完成课程 1。这是不可能的。
```

提示：

1. 输入的先决条件是由 边缘列表 表示的图形，而不是 邻接矩阵 。详情请参见[图的表示法](http://blog.csdn.net/woaidapaopao/article/details/51732947) 。
2. 你可以假定输入的先决条件中没有重复的边 。
3. `1 <= numCourses <= 10^5`

#### 思路
这道题可以用 DFS 来做，通过 DFS 来查找是否存在环。使用一个长度为 numCourses 的数组，如果未被访问则为 0，被当前的 DFS 访问过则为 1，被其他线路访问过则置为 -1；如果遇到 1，则直接返回 false，如果遇到 -1，则返回 true（因为之前访问过的线路没有返回 false，则可以保证后续不会存在 环）。

#### 伪代码
```
CAN-FINISH(numCourses, prerequisites)
    if numCourse == 0 || prerequisites.length == 0
        return true
    adjacency = new Map()
    flags = [numCourses]int{}
    for i = 0 to prerequisites.length -1
        adjacency[prerequisites[i][1]].add(prerequisites[i][0])
    for i = 0 to numCourses.length -1
        if !DFS(adjacency, flags, i)
            return false
    return true
        
DFS(adjacency, flags, course)
    if flags[course] == 1
        return false
    if flags[course] == -1
        return true
    flags[course] = 1
    for i = 0 to adjacency[course]
        if !DFS(adjacency, flags, adjacency[course][i])
            return false
    flags[course] = -1
    return true
```