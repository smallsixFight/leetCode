package flatten_nested_list_iterator

/*
来源：[leetCode](https://leetcode-cn.com/)
题目：[扁平化嵌套列表迭代器](https://leetcode-cn.com/problems/flatten-nested-list-iterator/)
标签：`栈` `设计`

简单思路：这题主要精力都花在看题并理解题目上了，有人在评论区说做题的人少，我个人看这题目看到一半都不太想做了。
题目看懂其实很简单，用递归处理 `NestedInteger` 就可以，这里不去写迭代的实现了。
另外，这道题没能提交测试案例，我不想去实现 `NestedInteger`（心态发生了微妙的变化0.0）。

官网运行结果记录
执行用时：4ms
内存消耗：5.6MB

*/

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type NestedInteger struct {
}

func (this NestedInteger) IsInteger() bool { return false }
func (this NestedInteger) GetInteger() int { return 1 }

func (this NestedInteger) GetList() []*NestedInteger { return nil }

type NestedIterator struct {
	data []int
}

func (this *NestedIterator) handleNested(list []*NestedInteger, res *[]int) {
	for i := range list {
		if list[i].IsInteger() {
			*res = append(*res, list[i].GetInteger())
		} else {
			this.handleNested(list[i].GetList(), res)
		}
	}
	return
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	iterator := NestedIterator{
		data: make([]int, 0),
	}
	iterator.handleNested(nestedList, &iterator.data)
	return &iterator
}

func (this *NestedIterator) Next() int {
	v := this.data[0]
	this.data = this.data[1:]
	return v
}

func (this *NestedIterator) HasNext() bool {
	return len(this.data) > 0
}
