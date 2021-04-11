package minimum_adjacent_swaps_for_k_consecutive_ones

import (
	"math"
)

/*
1703
来源：[leetCode](https://leetcode-cn.com/)
题目：[得到连续 K 个 1 的最少相邻交换次数](https://leetcode-cn.com/problems/minimum-adjacent-swaps-for-k-consecutive-ones/)
标签：`栈`

思路：这道题困难就困难在是一道数学题啊，先写下暴力法的思路（虽然超过时间限制了）：

- 首先将 nums 中所有为 1 的数字的位置整理出来生成一个新数组 arr；
- 然后以 arr 中每个元素为中心，分别向左右找到距离最小的位置放入连续的区间，直到达到 k 值；
- 扩展结束后，跟已经记录的结果值做比较，取较小值作为新结果值。

优化：
todo：看不懂公式，后续研究。

https://leetcode-cn.com/problems/minimum-adjacent-swaps-for-k-consecutive-ones/solution/de-dao-lian-xu-k-ge-1-de-zui-shao-xiang-lpa9i/

官网运行结果记录
执行用时：172ms(60%)/176ms(minMoves2)
内存消耗：8.8MB(80%)(minMoves2)

*/

func minMoves2(nums []int, k int) int {
	if k == 1 {
		return 0
	}
	g, total, count := make([]int, 0), []int{0}, 0
	for i := range nums {
		if nums[i] == 1 {
			g = append(g, i-count)
			total = append(total, total[len(total)-1]+g[len(g)-1])
			count++
		}
	}
	m, ans := len(g), math.MaxInt32
	for i := 0; i < m-k+1; i++ {
		mid := (i + i + k - 1) / 2
		q := g[mid]
		ans = min(ans, (2*(mid-i)-k+1)*q+(total[i+k]-total[mid+1])-(total[mid]-total[i]))
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minMoves(nums []int, k int) int {
	arr := make([]int, 0)
	for i := range nums {
		if nums[i] == 1 {
			arr = append(arr, i)
		}
	}
	res := math.MaxInt64
	for i := range arr {

		ln, rn := 0, 0
		prev, next := i-1, i+1
		temp := 0
		for ln+rn+1 < k {
			if next == len(arr) {
				temp += arr[i] - arr[prev] - ln - 1
				prev--
				ln++
			} else if prev == -1 {
				temp += arr[next] - arr[i] - rn - 1
				next++
				rn++
			} else if arr[i]-arr[prev] < arr[next]-arr[i] {
				temp += arr[i] - arr[prev] - ln - 1
				prev--
				ln++
			} else {
				temp += arr[next] - arr[i] - rn - 1
				next++
				rn++
			}
		}
		if temp < res {
			res = temp
		}
	}
	return res
}
