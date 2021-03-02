package group_anagrams

import "sort"

func GroupAnagramsThree(strs []string) [][]string {
	r := make(map[string][]string)
	for _, vv := range strs {
		a := []rune(vv)
		sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
		v := string(a)
		if _, ok := r[v]; !ok {
			r[v] = make([]string, 0)
		}
		r[v] = append(r[v], vv)
	}
	res := make([][]string, len(r))
	i := 0
	for idx := range r {
		res[i] = r[idx]
		i++
	}
	return res
}

// 空间优化
func GroupAnagramsTwoT(strs []string) [][]string {
	res := make([][]string, 0)
	prime := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103}
	record := make(map[int]int, 0)
	count := 0
	for idx := range strs {
		t := 1
		for _, x := range strs[idx] {
			t *= prime[x-'a']
		}
		if _, ok := record[t]; !ok {
			record[t] = count
			count++
			res = append(res, make([]string, 0))
		}
		res[record[t]] = append(res[record[t]], strs[idx])
	}
	return res
}

// 使用质数来计算
func GroupAnagramsTwo(strs []string) [][]string {
	res := make([][]string, 0)
	prime := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103}
	record := make(map[int][]string)
	for idx := range strs {
		t := 1
		for _, x := range strs[idx] {
			t *= prime[x-'a']
		}
		record[t] = append(record[t], strs[idx])
	}
	for idx := range record {
		res = append(res, record[idx])
	}
	return res
}

// 超时解法
func GroupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	next := strs
	for len(next) > 0 {
		val := next[0]
		sonRes := []string{val}
		tempNext := make([]string, 0)
		m := make(map[byte]uint8)
		if val != "" {
			for i := 0; i < len(val); i++ {
				m[val[i]]++
			}
		}
		for i := 1; i < len(next); i++ {
			if len(val) != len(next[i]) {
				tempNext = append(tempNext, next[i])
				continue
			}
			if next[i] == "" {
				if val == "" {
					sonRes = append(sonRes, next[i])
				} else {
					tempNext = append(tempNext, next[i])
				}
				continue
			} else if val == "" {
				tempNext = append(tempNext, next[i])
				continue
			} else {
				mm := make(map[byte]uint8)
				for k, v := range m {
					mm[k] = v
				}
				flag := true
				for j := 0; j < len(next[i]); j++ {
					if _, ok := mm[next[i][j]]; ok && mm[next[i][j]] > 0 {
						mm[next[i][j]]--
					} else {
						tempNext = append(tempNext, next[i])
						flag = false
						break
					}
				}
				if flag {
					sonRes = append(sonRes, next[i])
				}
			}
		}
		next = tempNext
		res = append(res, sonRes)
	}
	return res
}
