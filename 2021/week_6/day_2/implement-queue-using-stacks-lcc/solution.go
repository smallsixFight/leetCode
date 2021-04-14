package implement_queue_using_stacks_lcc

/*
面试题 03.04
来源：[leetCode](https://leetcode-cn.com/)
题目：[化栈为队](https://leetcode-cn.com/problems/implement-queue-using-stacks-lcci/)
标签：`栈` `设计`

思路：之前做过的题目。

官网运行结果记录
执行用时：0ms
内存消耗：1.9MB

*/

type MyQueue struct {
	s1 []int
	s2 []int
}

func Constructor() MyQueue {
	return MyQueue{s1: make([]int, 0), s2: make([]int, 0)}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.s2 = append(this.s2, x)
	this.s2 = append(this.s2, this.s1...)
	this.s1, this.s2 = this.s2, []int{}
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	v := this.s1[len(this.s1)-1]
	this.s1 = this.s1[:len(this.s1)-1]
	return v
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.s1[len(this.s1)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.s1) == 0
}
