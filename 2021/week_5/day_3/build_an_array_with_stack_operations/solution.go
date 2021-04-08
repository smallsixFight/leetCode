package build_an_array_with_stack_operations

/*
1441
来源：[leetCode](https://leetcode-cn.com/)
题目：[用栈操作构建数组](https://leetcode-cn.com/problems/build-an-array-with-stack-operations/)
标签：`栈`

思路：对题目进行解析后，可以发现结果集其实就是一个单调递增栈，使用一个指针 p 记录 target 中需要比较的数字，然后遍历 n 生成的数组，
当数组中的数字大于等于 target 当前的数字，则进行 push，否则，就是 push 和 pop 一起。

官网运行结果记录
执行用时：0ms
内存消耗：2.3MB

*/

func buildArray(target []int, n int) []string {
	p := 0
	res := make([]string, 0, n*2)
	for i := 1; i <= n && p < len(target); i++ {
		if target[p] == i {
			res = append(res, "Push")
			p++
		} else if target[p] > i {
			res = append(res, "Push", "Pop")
		}
	}
	return res
}
