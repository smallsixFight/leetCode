package maximal_rectangle

/*
来源：[leetCode](https://leetcode-cn.com/)
题目：[简化路径](https://leetcode-cn.com/problems/maximal-rectangle/)
标签：`栈` `数组` `哈希表` `动态规划`

简单思路：当尝试按层次遍历逐层添加高度并将图画出来后，可以发现每完成一层都需要求一次与 [柱状图中最大的矩阵](https://leetcode-cn.com/problems/largest-rectangle-in-histogram/)
这题中最大的矩形面积，然后与当前取得的最大面积比较。即求 matrix.length 次的largest-rectangle-in-histogram。


官网运行结果记录
执行用时：4ms
内存消耗：3.1MB

*/

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	maxArea := 0
	heights := make([]int, len(matrix[0]))
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == '0' {
				heights[j] = 0
			} else {
				heights[j] += 1
			}
		}
		maxArea = max(maxArea, largestRectangleArea(heights))
	}
	return maxArea
}

func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	maxArea := 0
	stack := stack{}
	for i := range heights {
		if stack.Len == 0 || heights[i] > heights[stack.peek()] {
			stack.push(i)
			continue
		}
		for stack.Len > 0 && heights[i] <= heights[stack.peek()] {
			idx := stack.pop()
			w := i
			if stack.Len > 0 {
				w = i - stack.peek() - 1 // -1 是因为不能包括 i
			}
			maxArea = max(maxArea, w*heights[idx])
		}
		stack.push(i)
	}
	for stack.Len > 0 {
		idx := stack.pop()
		// (idx - stack.peek()) + [(len(heights) - 1) - idx] // 左边宽度 + 右边宽度
		maxArea = max(maxArea, (len(heights)-1-stack.peek())*heights[idx])
	}
	return maxArea
}

type stack struct {
	data []int
	Len  int
}

func (s *stack) pop() int {
	l := len(s.data)
	if l == 0 {
		return 0
	}
	v := s.data[l-1]
	s.data = s.data[:l-1]
	s.Len--
	return v
}

func (s *stack) push(v int) {
	s.Len++
	s.data = append(s.data, v)
}

func (s *stack) peek() int {
	if s.Len == 0 {
		return -1
	}
	return s.data[s.Len-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
