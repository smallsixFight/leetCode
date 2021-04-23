package sort_list

import (
	"fmt"
	"sort"
)

/*
452
来源：[leetCode](https://leetcode-cn.com/)
题目：[ 用最少数量的箭引爆气球](https://leetcode-cn.com/problems/minimum-number-of-arrows-to-burst-balloons/)
标签：`排序` `贪心算法`

思路：大致处理逻辑如下：

- 按照 `points[i][0]，points[i][1]` 进行升序排序，然后遍历 points，每次遍历箭数加一；
- 然后使用内循环遍历后续的气球看能否被射爆；
- 使用 `r = points[i][1]` 作为初始值，内循环遍历 points，若 `points[j][0] <= r`，则表示可被一起射爆，然后 `r = min(r, points[j][1])`，缩小临界值，然后进行下回合比较；
- 不能遇到不能射爆的气球时结束内循环，进行下一次外循环。

官网运行结果记录
执行用时：88ms(67%)
内存消耗：7.2MB(67%)

*/

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] < points[j][1]
		}
		return points[i][0] < points[j][0]
	})
	fmt.Println(points)
	res := 0
	for i := 0; i < len(points); i++ {
		res++
		j := i + 1
		r := points[i][1]
		for j < len(points) && r >= points[j][0] {
			if r > points[j][1] {
				r = points[j][1]
			}
			j++
		}
		i = j - 1
	}
	return res
}
