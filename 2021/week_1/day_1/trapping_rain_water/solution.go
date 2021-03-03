package trapping_rain_water

/*
来源：[leetCode](https://leetcode-cn.com/)
题目：[接雨水](https://leetcode-cn.com/problems/trapping-rain-water/)
标签：`栈` `数组` `双指针` `动态规划`

简单思路：可知，需要两根具有高度且有距离的柱子才能接水，那么先从左到右遍历，使用双指针法累计数量，
但在题目图中的高度为 3 的柱子开始，计算就开始不准确了，因为后面没有高度大于等于 3 的柱子了。
接下来就考虑以最高柱子为中间点，分别计算其左侧与右侧能接的水，最后相加即可。这就符合动态规划，而最高的柱子就是
动态规划中的那个 m 值。
此外，左右两个不用再找最高的柱子了，因为已经确定左右两侧的最后的柱子一定是最高的，那么再结合前面想到的双指针法
就可以解了。

官网运行结果记录
结果：通过
执行用时：0ms/4ms
内存消耗：2.8MB

*/

func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	m := 0
	for i := 1; i < len(height); i++ {
		if height[i] > height[m] {
			m = i
		}
	}
	s, f := 0, 1
	total, tempTotal := 0, 0
	for f <= m {
		if height[s] > height[f] {
			tempTotal += height[s] - height[f]
		} else {
			total += tempTotal
			s = f
			tempTotal = 0
		}
		f++
	}
	s = len(height) - 1
	f = s - 1
	for f >= m {
		if height[s] > height[f] {
			tempTotal += height[s] - height[f]
		} else {
			total += tempTotal
			s = f
			tempTotal = 0
		}
		f--
	}
	return total
}
