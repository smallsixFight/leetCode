package count_and_say

import (
	"strings"
)

func CountAndSay2(n int) string {
	arr := make([]string, n)
	arr[0] = "1"
	for i := 1; i < n; i++ {
		if arr[i] != "" {
			return arr[i]
		}
		count := 0
		last := arr[i-1]
		result := strings.Builder{}
		for p1, p2 := 0, 0; p2 < len(last); p2++ {
			if last[p1] == last[p2] {
				count++
				if p2 == len(last)-1 {
					result.WriteString(string('0' + count))
					result.WriteByte(last[p1])
				}
			} else {
				result.WriteString(string('0' + count))
				result.WriteByte(last[p1])
				count = 0
				p1 = p2
				p2--
			}
		}
		arr[i] = result.String()
	}
	return arr[n-1]
}

func CountAndSay(n int) string {
	arr := make([]string, n)
	arr[0] = "1"
	for i := 1; i < n; i++ {
		if arr[i] != "" {
			return arr[i]
		}
		count := 0
		p1, p2 := 0, 0
		last := arr[i-1]
		result := strings.Builder{}
		for p2 < len(last) {
			if last[p1] == last[p2] {
				count += 1
				p2 += 1
				if p2 == len(last) {
					result.WriteString(string('0' + count))
					result.WriteByte(last[p1])
				}
			} else {
				result.WriteString(string('0' + count))
				result.WriteByte(last[p1])
				count = 0
				p1 = p2
			}
		}
		arr[i] = result.String()
	}
	return arr[n-1]
}
