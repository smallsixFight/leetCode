package permutation_sequence

import (
	"sort"
)

func GetPermutationTwo(n int, k int) string {
	if n == 1 {
		return "1"
	}
	record := make(map[byte]bool)
	res := make([]byte, n)
	num := 1 // 当前组合具备的行数
	for i := n; i >= 1; i-- {
		num *= i
	}
	l := n
	var min uint8 = '1'
	for i := 0; i < n; i++ {
		num /= l // 当前每个数字在第一位的组合的个数
		val, add := min, (k-1)/num
		for add > 0 {
			val += 1
			if !record[val] { // 如果当前数字已被使用，不计数
				add--
			}
		}
		for record[val] {
			val += 1
		}
		record[val] = true
		for record[min] {
			min++
		}
		res[i] = val
		k -= (k - 1) / num * num
		l--
	}
	return string(res)
}

func GetPermutation(n int, k int) string {
	if n == 1 {
		return "1"
	}
	res := make([]byte, n)
	num := 1
	for i := n; i > 1; i-- {
		num *= i
	}
	rowNum := num / n                 // 倍数
	res[0] = byte('1' + (k-1)/rowNum) // 首个数字
	for i, j := 1, 1; j < len(res); i++ {
		if uint8(i) == res[0]-'0' {
			continue
		}
		res[j] = '0' + uint8(i)
		j++
	}
	row := k - (k-1)/rowNum*rowNum // 所在行数
	// 迭代获取下一个最大值
	for a := 1; a < row; a++ {
		p1, p2 := -1, 0
		for i := len(res) - 1; i > 1; i-- {
			for j := i - 1; j > 0; j-- {
				if res[i] > res[j] && j > p1 {
					p1 = j
					p2 = i
					break
				}
			}
		}
		if p1 != -1 {
			res[p1], res[p2] = res[p2], res[p1]
			arr := res[p1+1:]
			sort.Slice(arr, func(i, j int) bool {
				return arr[i] < arr[j]
			})
		}
	}
	return string(res)
}
