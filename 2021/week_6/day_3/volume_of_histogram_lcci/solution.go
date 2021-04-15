package volume_of_histogram_lcci

/*
面试题 17.21
来源：[leetCode](https://leetcode-cn.com/)
题目：[直方图的水量](https://leetcode-cn.com/problems/volume-of-histogram-lcci/)
标签：`栈` `数组` `双指针`

思路：这道题是第一周做过的一道题，这里重新整理一遍思路：

- 从左向右遍历，先找到最高的柱子的其中一个；
- 使用双指针法，分别从左向右和从右向左遍历，以找到的最高柱子为终点；
- 双指针的做法是，快指针（f）指向的柱子若比慢指针（s）指向的柱子低，则计算高度并累计，若 f 指向的柱子高度大于等于 s 指向的柱子的高度，将 s 移动到 f。以此类推，直至遍历结束。

官网运行结果记录
执行用时：4ms
内存消耗：2.8MB

*/

func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	total, m := 0, 0
	for i := 1; i < len(height); i++ {
		if height[i] > height[m] {
			m = i
		}
	}
	s, f := 0, 1
	for f <= m {
		if height[f] < height[s] {
			total += height[s] - height[f]
		} else {
			s = f
		}
		f++
	}
	s, f = len(height)-1, len(height)-2
	for f >= m {
		if height[f] < height[s] {
			total += height[s] - height[f]
		} else {
			s = f
		}
		f--
	}
	return total
}
