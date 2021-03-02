package valid_parentheses

import (
	"github.com/leetCode/antiquated/utils"
)

func IsValid2(s string) bool {
	if s == "" {
		return true
	}
	sli := make([]byte, 0)
	point := -1
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			point++
			sli = append(sli, s[i])
			if len(sli) > len(s[i:]) {
				return false
			}
		} else {
			if len(sli) == 0 {
				return false
			}
			val := sli[point]
			sli = sli[0:point]
			point--
			if s[i] == ')' && val != byte('(') {
				return false
			}
			if s[i] == ']' && val != byte('[') {
				return false
			}
			if s[i] == '}' && val != byte('{') {
				return false
			}
		}
	}
	return len(sli) == 0
}

func IsValid(s string) bool {
	if s == "" {
		return true
	}
	stack := utils.NewStack(len(s) / 2)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack.Push(s[i])
			if stack.Len() > len(s[i:]) {
				return false
			}
		} else {
			if stack.IsEmpty() {
				return false
			}
			val, _ := stack.Pop()
			if s[i] == ')' && val != byte('(') {
				return false
			}
			if s[i] == ']' && val != byte('[') {
				return false
			}
			if s[i] == '}' && val != byte('{') {
				return false
			}
		}
	}
	return stack.IsEmpty()
}
