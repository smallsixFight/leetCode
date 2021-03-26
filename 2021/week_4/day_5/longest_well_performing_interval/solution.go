package longest_well_performing_interval

/*
1124
来源：[leetCode](https://leetcode-cn.com/)
题目：[表现良好的最长时间段](https://leetcode-cn.com/problems/longest-well-performing-interval/)
标签：`栈`

简单思路：这道题先把时间划分为疲劳和非疲劳两种类型，将 hours 的值根据前面说的类型分别替换为 1 和 -1。那么这道题就是找到和大于 1 的最长子字符串。
先从左向右遍历 hours，生成一个前 n 项和的数组，然后通过两两互减来获取符合条件的最长字符串。

前面的思路提交后，耗时长达 788ms，只打败了 7% 的答题者，然后通过看题解，了解到用单调栈来减少无用的计算比较，大致思路整理如下：
可以知道，前 n 项和中值越大的子字符串是符合答案的最长子字符串的有力竞争者，反过来想，值越小的表示是连续不符合要求的最大子字符串，应该移除，那么可以使用一个单调递减栈结构来生成一个可以用于剔除获取符合要求的子字符串。

官网运行结果记录
执行用时：788ms(7%)	24ms(96%)
内存消耗：6.1MB		6.4MB(62%)

*/

func longestWPI2(hours []int) int {
	for i := range hours {
		if hours[i] > 8 {
			hours[i] = 1
		} else {
			hours[i] = -1
		}
	}
	record := make([]int, len(hours)+1)
	record[0] = 0
	for i := 0; i < len(hours); i++ {
		record[i+1] = record[i] + hours[i]
	}
	stack := make([]int, 0)
	for i := range record {
		if len(stack) == 0 || record[i] < record[stack[len(stack)-1]] {
			stack = append(stack, i)
		}
	}
	res := 0
	for i := len(record) - 1; i >= 0; i-- {
		for len(stack) > 0 && record[stack[len(stack)-1]] < record[i] {
			if i-stack[len(stack)-1] > res {
				res = i - stack[len(stack)-1]
			}
			stack = stack[:len(stack)-1]
		}
	}
	return res
}

func longestWPI(hours []int) int {
	for i := range hours {
		if hours[i] > 8 {
			hours[i] = 1
		} else {
			hours[i] = -1
		}
	}
	record := make([]int, len(hours)+1)
	record[0] = 0
	for i := 0; i < len(hours); i++ {
		record[i+1] = record[i] + hours[i]
	}
	//fmt.Println(record)
	res := 0
	for i := len(record) - 1; i > 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if record[i]-record[j] > 0 && i-j > res {
				res = i - j
			}
		}
	}
	if res == 0 && hours[0] == 1 {
		res = 1
	}
	return res
}
