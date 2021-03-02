### 课程表 II（Course Schedule II）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/course-schedule-ii/)

#### 题目
现在你总共有 n 门课需要选，记为`0`到`n-1`。

在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们:`[0,1]`

给定课程总量以及它们的先决条件，返回你为了学完所有课程所安排的学习顺序。

可能会有多个正确的顺序，你只要返回一种就可以了。如果不可能完成所有课程，返回一个空数组。

示例 1：
```
输入: 2, [[1,0]] 
输出: [0,1]
解释: 总共有 2 门课程。要学习课程 1，你需要先完成课程 0。因此，正确的课程顺序为 [0,1] 。
```
示例 2：
```
输入: 4, [[1,0],[2,0],[3,1],[3,2]]
输出: [0,1,2,3] or [0,2,1,3]
解释: 总共有 4 门课程。要学习课程 3，你应该先完成课程 1 和课程 2。并且课程 1 和课程 2 都应该排在课程 0 之后。
     因此，一个正确的课程顺序是 [0,1,2,3] 。另一个正确的排序是 [0,2,1,3] 。
```

说明：

1. 输入的先决条件是由边缘列表表示的图形，而不是邻接矩阵。详情请参见图的表示法。
2. 你可以假定输入的先决条件中没有重复的边。

提示：
1. 这个问题相当于查找一个循环是否存在于有向图中。如果存在循环，则不存在拓扑排序，因此不可能选取所有课程进行学习。
2. 通过 DFS 进行拓扑排序 - 一个关于Coursera的精彩视频教程（21分钟），介绍拓扑排序的基本概念。
3. 拓扑排序也可以通过 BFS 完成。

#### 思路
这道题跟前面做的`课程表`区别不大，这道题我依旧使用 DFS 来做，然后由于学习某个课程需要将所有前置课程学习完，所以我使用了一个栈的结构来存储结果集。

#### 伪代码
```
FIND-ORDER(numCourses, prerequisites)
    if numCourses == 0
        return []int{}
    adjacency = new Map()
    flags = [numCourses]int{}
    res = [numCourses]int{}
    idx = numCourses -1
    for i = 0 to prerequisites.length -1
        adjacency[prerequisites[i][1]].add(prerequisites[i][0])
    for i = 0 to numCourses.length -1
        if !DFS(adjacency, flags, i, res)
            return []int{}
    retur res

DFS(adjacency, flags course, res)
    if flags[course] == 1
        return false
    if flags[course] == -1
        return true
    flags[course] = 1
    for i = 0 to adjacency[course].length -1
        if !DFS(adjacency, flags adjacency[course][i], res)
            return false
    flags[course] = -1
    res[idx] = course
    idx --
    return false
```