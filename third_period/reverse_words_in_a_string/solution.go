package reverse_words_in_a_string

func ReverseWords(s string) string {
	if len(s) == 0 {
		return s
	}
	res := make([]byte, 0)
	p1, p2 := len(s)-1, len(s)-1
	for ; p1 >= 0; p1-- {
		if s[p2] == ' ' {
			p2--
		} else if s[p1] == ' ' {
			res = append(res, []byte(s[p1+1:p2+1]+" ")...)
			p2 = p1
		}
	}
	if s[0] != ' ' {
		res = append(res, []byte(s[p1+1:p2+1]+" ")...)
	}
	if len(res) > 0 {
		return string(res[:len(res)-1])
	}
	return string(res)
}
