package verify_preorder_serialization_of_a_binary_tree

import "strings"

/*
来源：[leetCode](https://leetcode-cn.com/)
题目：[验证二叉树的前序序列化](https://leetcode-cn.com/problems/verify-preorder-serialization-of-a-binary-tree/)
标签：`栈`

简单思路
先整理前序遍历序列化会满足的特性：
- 由于空结点会用 ‘#’ 代替，所以每个不为空的结点都会有两个子结点，序列化后结点数量一定是奇数，加上逗号的长度的字符串的长度一定是偶数（除非根结点为 null）。
- 将序列化后的字符串使用一个栈做出入栈操作：遇到不为空的结点进行入栈，该结点的两个子结点都处理完成后则该结点可以出栈，遍历完成后栈肯定是空的。

通过上面的特性进行反推，可以得出不满足前序遍历序列化的条件：
- 第一个元素为 `#` 但字符串长度大于 1；
- 栈为空却要进行出栈操作；
- 遍历完后栈不为空。

官网运行结果记录 ()
执行用时：0ms/4ms
内存消耗：3.2MB

------

由于内存消耗对比下过大，于是去看了题解，学到了 isValidSerialization2 的写法，思路整理如下：
使用 slot 的概念，每个结点都需要消耗一个 slot，如果一个结点为非空结点，则需要为它的两个子结点添加 slot。如果没有遍历结束，slot 为空则返回 false，遍历结束后，正确的前序遍历序列化的 slot 应该为 0。

官网运行结果记录（isValidSerialization2）
执行用时：0ms
内存消耗：2.7MB

官网运行结果记录（isValidSerialization3）
执行用时：0ms
内存消耗：2.1MB

*/

func isValidSerialization(preorder string) bool {
	if len(preorder) == 0 || (preorder[0] == '#' && len(preorder) > 1) ||
		(len(preorder) == 1 && preorder[0] != '#') {
		return false
	}
	if len(preorder) == 1 && preorder[0] == '#' {
		return true
	}

	arr := strings.Split(preorder, ",")
	if len(arr)%2 == 0 {
		return false
	}
	stack, record := make([]int, 0, len(arr)), make([]int, len(arr))
	for i := range arr {
		if i != 0 && len(stack) == 0 {
			return false
		}
		if arr[i] == "#" {
			if len(stack) == 0 {
				return false
			}
			record[stack[len(stack)-1]]++
			for len(stack) > 0 && record[stack[len(stack)-1]] == 2 {
				stack = stack[:len(stack)-1]
				if len(stack) > 0 {
					record[stack[len(stack)-1]]++
				}
			}
		} else {
			stack = append(stack, i)
		}
	}

	return len(stack) == 0
}

func isValidSerialization2(preorder string) bool {
	slots := 1
	arr := strings.Split(preorder, ",")
	for i := range arr {
		if slots == 0 {
			return false
		}
		if arr[i] == "#" {
			slots--
		} else {
			slots++
		}
	}
	return slots == 0
}

func isValidSerialization3(preorder string) bool {
	slots := 1
	for i := 0; i < len(preorder); {
		if slots == 0 {
			return false
		}
		if preorder[i] == ',' {
			i++
		}
		if preorder[i] == '#' {
			slots--
			i++
		} else {
			slots++ // slot = slot -1 + 2
			for i < len(preorder) && preorder[i] != ',' {
				i++
			}
		}
	}
	return slots == 0
}
