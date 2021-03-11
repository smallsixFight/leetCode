package mini_parser

/*
来源：[leetCode](https://leetcode-cn.com/)
题目：[扁平化嵌套列表迭代器](https://leetcode-cn.com/problems/mini-parser/)
标签：`栈` `字符串`

简单思路：使用一个栈用来保存 NestedInteger，从头向尾遍历，左括号使用 nil 替换，数字存为一个 NestedInteger，遇到右括号时，进行出栈直到遇到第一个 nil，创建一个 NestedInteger，然后 Add 出栈的 NestedInteger，然后将新的 NestedInteger 入栈，以此直到遍历结束。最后取出栈顶的 NestedInteger 进行返回。

官网运行结果记录
执行用时：0ms
内存消耗：5.1MB

*/

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */

type NestedInteger struct {
}

func (n *NestedInteger) SetInteger(value int)   {}
func (n *NestedInteger) Add(elem NestedInteger) {}

func deserialize(s string) *NestedInteger {
	if len(s) == 0 {
		return nil
	}
	stack := make([]*NestedInteger, 0)
	leftRecord := make([]int, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == ',' {
			continue
		}
		if s[i] == '[' {
			leftRecord = append(leftRecord, len(stack))
			stack = append(stack, nil)
		} else if s[i] == ']' {
			n := NestedInteger{}
			j := leftRecord[len(leftRecord)-1]
			for k := j + 1; k < len(stack); k++ {
				n.Add(*stack[k])
			}
			stack = stack[:j]
			stack = append(stack, &n)
			leftRecord = leftRecord[:len(leftRecord)-1]
		} else {
			num := 0
			direct := 1
			j := i
			if s[i] == '-' {
				direct = -1
				j++
			}
			for ; j < len(s) && s[j] != ',' && s[j] != ']'; j++ {
				num = num*10 + int(s[j]-'0')
			}
			i = j - 1
			n := NestedInteger{}
			n.SetInteger(num * direct)
			stack = append(stack, &n)
		}
	}
	return stack[0]
}
