package next_greater_element_i

/*
来源：[leetCode](https://leetcode-cn.com/)
题目：[下一个更大元素 I](https://leetcode-cn.com/problems/next-greater-element-i/)
标签：`栈`

简单思路：使用一个 hashTable 记录 num2 中每个数字中每个数字的下一个更大数字，然后遍历 num1 通过 hashTable 获取结果后返回。
处理 num2：从左向右遍历 num2，依次入栈，如果新的元素比栈顶元素大，则以为该元素是栈顶元素后第一个更大的元素，进行记录并出栈，内循环前面的比较直到栈顶元素大于新元素，最后将新元素入栈。
遍历结束后，若栈不为空，依次出栈，并记录为 -1。

官网运行结果记录
执行用时：4ms
内存消耗：3MB

*/

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	record := make(map[int]int, len(nums2))
	stack := make([]int, 0, len(nums1))
	for i := range nums2 {
		for len(stack) > 0 && nums2[i] > stack[len(stack)-1] {
			v := stack[len(stack)-1]
			record[v] = nums2[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums2[i])
	}
	for len(stack) > 0 {
		record[stack[len(stack)-1]] = -1
		stack = stack[:len(stack)-1]
	}
	for i := range nums1 {
		nums1[i] = record[nums1[i]]
	}
	return nums1
}
