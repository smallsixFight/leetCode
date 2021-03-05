package largest_rectangle_in_histogram

/*
来源：[leetCode](https://leetcode-cn.com/)
题目：[柱状图中最大的矩形](https://leetcode-cn.com/problems/largest-rectangle-in-histogram/)
标签：`栈` `数组`

简单思路：通过观察，最矮的柱子具有最大的宽度，以其为分割点，其左右两侧最矮的柱子又有各自那一侧最大的宽度；
另外，最矮的柱子不需要顾虑比它高的柱子，它能取得的最大面积就是它的高度 * 最大宽度。
但总不能每次去找最矮的柱子进行分割，那么逆向一下，最矮的柱子的宽度一定会包含高度大于等于它的柱子，而面积的计算
需要向最低高度的柱子看齐，那么从左向右遍历，以当前的柱子为最低高度，向左右遍历直到遇到高度更低的柱子，那么就能
得到宽度计算面积。因为按顺序遍历，那么左侧较矮的柱子计算的最大面积肯定包含当前柱子。


官网运行结果记录（使用 largestRectangleArea2）
执行用时：96ms/100ms/104ms
内存消耗：8MB/8.1MB

*/

// 思路的暴力解法提交后超时间限制，通过题解学到单调栈的用法。
// 在原来的思路中，需要每根柱子都需要向左右遍历找到边界（比当前柱子矮的柱子），但在一大堆相同高度的柱子中
// 进行计算时，就会有很多重复的无用计算，比如让 largestRectangleArea 超时的第 93 个测试案例。
// 不过，暴力解法也是单调栈解法的基础，再更深地思考，就能想通单调栈解法。
// 从左向右遍历，第一根柱子左侧没有比它还要矮的柱子，所以只要向右寻找边界，所以右侧是比它高的柱子的话，右侧的
// 柱子也要寻找边界，而新柱子左侧边界就是前一根矮柱子，另外，新柱子如果遇到边界，则宽度确定，而旧柱子还有可能
// 需要继续向后寻找边界，也有可能也是它的边界。
// 根据前面的描述，总是比当前柱子高的柱子先得出面积，有后进先出的模式，那么就可以用栈来处理。
// 先将第一根柱子的索引入栈，然后向后遍历，当遇到的柱子高度大于栈顶的柱子，则新的柱子入栈；当遇到的柱子高度小于等于
// 栈顶的柱子的高度，则意味栈顶的柱子找到了边界，可以出栈求面积，然后将新的柱子入栈。
// 遍历完成后，栈里如果还有柱子，则栈里的柱子的右边界是最后的柱子，全部出栈求面积。

// largestRectangleArea2 仅是在粗略看下官方题解找到的思路，所以答案并不一致，且 largestRectangleArea2 貌似
// 与官方运行的答案对比，内存占用少了 0.5MB 左右。
func largestRectangleArea2(heights []int) int {
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

func largestRectangleArea(heights []int) int {
	maxArea := 0
	for i := 0; i < len(heights); i++ {
		h := heights[i]
		w := 1
		for j := i - 1; j >= 0; j-- {
			if heights[j] < heights[i] {
				break
			}
			w++
		}
		for j := i + 1; j < len(heights); j++ {
			if heights[j] < heights[i] {
				break
			}
			w++
		}
		maxArea = max(maxArea, h*w)
	}
	return maxArea
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
