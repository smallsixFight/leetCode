package insert_interval

/*
57
来源：[leetCode](https://leetcode-cn.com/)
题目：[插入区间](https://leetcode-cn.com/problems/insert-interval/)
标签：`排序` `数组`

简单思路：使用一个 res 保存结果集，从头开始遍历 按照下面的逻辑处理：

- 若已经执行了插入操作或 `newInterval[0]>intervals[i][1]`，`intervals[i][1]` 直接放入结果集；
- 若 `newInterval[1]<intervals[i][0]`，意味不能合并，直接插入（这里其实隐藏着 `newInterval[1] < intervals[i][0]`）；
- 创建一个新数组，起始位置取 `min(min(newInterval[0], intervals[i][1]), intervals[i][0])`，结束位置取 `max(newInterval[1], intervals[i][1])`，然后使用内存向后遍历尝试合并，直至不能合并为止，最后新数组放入结果集；

官网运行结果记录
执行用时：4ms
内存消耗：4.6MB

*/

func insert(intervals [][]int, newInterval []int) [][]int {
	res, isInsert := make([][]int, 0), false
	for i := 0; i < len(intervals); i++ {
		if isInsert || newInterval[0] > intervals[i][1] {
			res = append(res, intervals[i])
		} else {
			if newInterval[1] < intervals[i][0] { // 不能合并，直接插入
				res = append(res, newInterval, intervals[i])
			} else { // 可进行合并
				item := []int{min(min(newInterval[0], intervals[i][1]), intervals[i][0]), max(newInterval[1], intervals[i][1])}
				j := i + 1
				for j < len(intervals) && item[1] >= intervals[j][0] {
					item[1] = max(intervals[j][1], item[1])
					j++
				}
				i = j - 1
				res = append(res, item)
			}
			isInsert = true
		}
	}
	if !isInsert {
		res = append(res, newInterval)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
