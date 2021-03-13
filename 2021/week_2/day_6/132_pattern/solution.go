package _32_pattern

/*
来源：[leetCode](https://leetcode-cn.com/)
题目：[132 模式](https://leetcode-cn.com/problems/132-pattern/)
标签：`栈`

思路：

- 从左向右遍历，记录当前位置左侧的最小数字；
- 使用一个单调递减栈来保存可用来进行比较的 a[k] 数字；
- 从右向左遍历，在 a[i] 和 a[j] 满足 `13` 模式（当作 `12` 模式来判断就行）的情况下，进行下面的处理：
    - 当栈顶的数字小于等于当前的数字 a[i] （即从左到 j 为止最小的数字） 时， 表示 a[k] 无法与之前的数字构成 `132` 模式，可以直接抛弃，依次出栈直到栈为空亦或栈顶元素大于 a[i]；
    - 当栈顶的数字小于当前的数字 a[j] 时，再将栈顶数字与 a[i] 比较，如果大于 a[i]，则满足 `132` 模式，可直接返回 true；
    - 最后如果不能返回 true，将 a[j] 入栈。
- 关于可以再 a[i] 和 a[j] 满足 `13` 模式进行处理，不满足的可以忽略的说明：这是由于 a[i] 其实是 a[j] 左侧最小的数字，如果不满足 `13` 模式，则表示 a[j] = a[i]，a[j] 与左侧的任何数字不都能构成 `13` 模式，那么也不嫩作为 a[k] 的参数，这是由于当前数字从左到当前位置最小的数字，那么也就不能满足 a[k] > a[i] 的条件，所以可以忽略跳过。


官网运行结果记录
执行用时：8ms/12ms
内存消耗：5.2MB

*/

func find132pattern(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	minRecord, stack := make([]int, len(nums)), make([]int, 0)
	minRecord[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		minRecord[i] = min(nums[i], minRecord[i-1])
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] > minRecord[i] {
			for len(stack) > 0 && minRecord[i] >= stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			}
			// len(stack) > 0 表示 存在 `132` 模式中的 `12`
			if len(stack) > 0 && stack[len(stack)-1] < nums[i] {
				return true
			}
			stack = append(stack, nums[i])
		}
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
