package yong_liang_ge_zhan_shi_xian_dui_lie_lcof

/*
剑指 Offer 09
来源：[leetCode](https://leetcode-cn.com/)
题目：[用两个栈实现队列](https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/)
标签：`栈` `设计`

思路：之前好像有做过用栈实现队列的题目来着。用一个 slice 来实现就好。

官网运行结果记录
执行用时：232ms(93%)
内存消耗：8.5MB(62%)

*/

type CQueue struct {
	data []int
}

func Constructor() CQueue {
	return CQueue{data: make([]int, 0)}
}

func (this *CQueue) AppendTail(value int) {
	this.data = append(this.data, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.data) > 0 {
		v := this.data[0]
		this.data = this.data[1:]
		return v
	}
	return -1
}
