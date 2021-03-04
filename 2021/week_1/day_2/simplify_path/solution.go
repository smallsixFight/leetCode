package simplify_path

import (
	"strings"
)

/*
来源：[leetCode](https://leetcode-cn.com/)
题目：[简化路径](https://leetcode-cn.com/problems/simplify-path/)
标签：`栈` `字符串`

简单思路：从左到右遍历，将每个有效的路径放到栈中，遍历完成后取出组合为结果。如果最后栈没有数据，返回根路径（`/`)。

遍历判断逻辑：
1. 使用一个 []byte，默认第一个元素为 `/`，开始遍历；
2. 遇到 `/` ，如果 []byte 长度为 1 则 continue；如果 []byte 长度不为 1：
	1. 如果 `.` 的标记存在并且 []byte 长度为 2，则以为是 `/./`，那么忽略并重置 []byte；
	2. 否则 []byte 是一个有效路径，进行 push并重置 []byte；
3. 遇到 `.`，如没做标记则做标记，并放入 []byte；如果标记存在，那么当下一位的值是 `/` 且 []byte 长度为 2 时，则
将栈进行 pop并重置 []byte；
4. 遇到其他的值，则直接放入 []byte 即可。

犯的错误：
- 没有处理最后一个符号是 `/` 的情况。

官网运行结果记录
执行用时：0ms/4ms
内存消耗：2.9MB/3.1MB：将 `tempPath = []byte{'/'}` 改为 `tempPath = tempPath[:1]` 降低了内存消耗。

*/

// 用 slice 模拟栈结构
type stack struct {
	data []string
	Len  int
}

func (s *stack) pop() string {
	l := len(s.data)
	if l == 0 {
		return ""
	}
	v := s.data[l-1]
	s.data = s.data[:l-1]
	s.Len--
	return v
}

func (s *stack) push(v string) {
	s.Len++
	s.data = append(s.data, v)
}

func (s *stack) String() string {
	sb := strings.Builder{}
	for i := range s.data {
		sb.WriteString(s.data[i])
	}
	return sb.String()
}

func simplifyPath(path string) string {
	tempPath := []byte{'/'}
	pointFlag := false
	stack := stack{}
	for i := 0; i < len(path); i++ {
		if path[i] == '/' {
			if len(tempPath) == 1 || i == len(path)-1 {
				continue
			}
			if !pointFlag {
				stack.push(string(tempPath))
			} else {
				pointFlag = false
			}
			tempPath = tempPath[:1]
		} else if path[i] == '.' {
			if len(tempPath) == 1 {
				pointFlag = true
			} else if pointFlag && len(tempPath) == 2 && i+1 < len(path) && path[i+1] == '/' {
				stack.pop()
				tempPath = tempPath[:1]
				continue
			}
			tempPath = append(tempPath, path[i])
		} else {
			tempPath = append(tempPath, path[i])
			pointFlag = false
		}
	}
	if len(tempPath) > 1 {
		v := string(tempPath)
		if v == "/.." {
			stack.pop()
		} else if v != "/." {
			stack.push(v)
		}
	}
	if stack.Len == 0 {
		return "/"
	}
	return stack.String()
}
