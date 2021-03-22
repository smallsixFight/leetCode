package sum_of_subarray_minimums

/*
来源：[leetCode](https://leetcode-cn.com/)
题目：[子数组的最小值之和](https://leetcode-cn.com/problems/sum-of-subarray-minimums/)
标签：`栈` `数组`

简单思路：这道题其实跟 `柱状图中的最大矩形` 是一样的，只是需要做最大面的汇总，而且矩阵那道题有张图，更容易去想。这里就当作复习吧，思路整理如下：
从左向右遍历，以当前的位置向左右两侧扩展遍历，左侧以不大于当前值作为边界，左侧以小于当前值作为边界，计算并汇总。

上面的思路提交后超出时间限制，但是思路是对的，于是顺着尝试进行优化。观察原来的写法，可以发现，每次都会进行中心扩展，但后续的扩展可能会重复之前的扩展比较，而且会很多次，这都是无意义的，那么可以先把查找当前位置的数字的左右两侧边界分别先查出来，利用单调栈来实现，再进行计算。
计算方面，对上面思路的统计方法进行整理，可以发现中心结点左侧的子集数为中心与左侧边界的距离加一，即 `i-left+1`，右侧的子集数除了与前面左侧子集一样的方法计算数量，即 `right-i`，还有组合左侧的每个子集作为新子集，那么可以得到公式 `(i-left+1) + (right-i) + (i-left) * (right-i) => (i-left+1) * (right-i+1)`

官网运行结果记录
执行用时：68ms(78%)/72ms(46%)
内存消耗：7.8MB(7%)

*/

func sumSubarrayMins2(arr []int) int {
	prev, next := make([]int, len(arr)), make([]int, len(arr))
	stack := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		for len(stack) > 0 && arr[i] <= arr[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		prev[i] = -1
		if len(stack) > 0 {
			prev[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	stack = make([]int, 0, len(stack))
	for i := len(arr) - 1; i >= 0; i-- {
		for len(stack) > 0 && arr[i] < arr[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		next[i] = len(arr)
		if len(stack) > 0 {
			next[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	var res int64
	for i := range arr {
		res += int64((i-prev[i])*(next[i]-i)%(1e9+7)) * int64(arr[i]%(1e9+7))
		res %= (1e9 + 7)
	}
	return int(res)
}

func sumSubarrayMins(arr []int) int {
	res := 0
	for i := range arr {
		left, total := i-1, arr[i]
		for ; left >= 0 && arr[left] > arr[i]; left-- {
			total += arr[i] // 1 1 1	3
		}
		for j := i + 1; j < len(arr) && arr[j] >= arr[i]; j++ {
			// 2 3 4	3
			total += (i - left) * (arr[i] % (1e9 + 7)) // 需要连续的子数组
		}
		res += total % (1e9 + 7)
	}
	return res % (1e9 + 7)
}
