package decoded_string_at_index

/*
来源：[leetCode](https://leetcode-cn.com/)
题目：[索引处的解码字符串](https://leetcode-cn.com/problems/decoded-string-at-index/)
标签：`栈`

简单思路：这道题使用一个栈来保存解码后放入磁带中的字符，另外整个字符串 S 其实并不需要完全解析，只需要解析达到索引 K 的位置即可，大致处理逻辑如下：

- 使用一个栈 s1 作为存储，从左向右遍历；
- 如果是字符，直接入栈，如果是数字，取出栈内所有元素，重复入栈 `num` 次（注意，这里不减一是因为栈内元素全部出栈了，需要重新再入）；
- 如果 `s1.size + repeatStr.length >= K`，那么进行 `n = K - s1.size`，然后遍历 `repeatStr`，找到位于 n 索引的字符返回即可。

这个思路直接内存溢出，放弃。

思路二：通过题解的逆向法，重新思考并进行了整理：

- 如果一个示例为 `apple6 24`，可以知道的是，返回的值肯定在 `apple` 之中，我们进行 `n = 24 % len(apple)` 运算后，返回值就是 str[n-1]；
- 如果一个示例为 `apple26 24`，的计算也是同理，只要遇到数字直接将已经得到的解码字符串的长度除以这个数字即可；
- 解码后的字符串的长度只需要从左向右遍历获取即可。另外，可以不用像题解那样求出完整解码后的字符串的长度，只要找到解码后字符串长度大于等于 K 的位置即可。

官网运行结果记录
执行用时：0ms
内存消耗：1.9MB

*/

func decodeAtIndex(S string, K int) string {
	p, size := 0, 0
	for ; p < len(S); p++ {
		if 'a' <= S[p] && S[p] <= 'z' {
			size++
		} else {
			size *= int(S[p] - '0')
		}
		if size >= K {
			break
		}
	}
	for ; p >= 0; p-- {
		K %= size
		if K == 0 && 'a' <= S[p] && S[p] <= 'z' {
			return string(S[p])
		}
		if 'a' <= S[p] && S[p] <= 'z' {
			size--
		} else {
			size /= int(S[p] - '0')
		}
	}
	return ""
}

// deprecated
func decodeAtIndex2(S string, K int) string {
	s1 := make([]byte, 0)
	for i := range S {
		if 'a' <= S[i] && S[i] <= 'z' {
			s1 = append(s1, S[i])
		} else {
			repeatStr, num := string(s1), int(S[i]-'0')
			for j := 1; j < num; j++ {
				if len(s1) >= K {
					return string(s1[K-1])
				} else if len(s1)+len(repeatStr) >= K {
					return string(repeatStr[K-len(s1)-1])
				} else {
					s1 = append(s1, repeatStr...)
				}
			}

		}
	}
	if len(s1) != 0 {
		return string(s1[K-1])
	}
	return ""
}
