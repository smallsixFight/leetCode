package letter_combinations_of_a_phone_number

func Solution3(digits string) []string {
	res := make([]string, 0)
	if len(digits) < 1 || digits[0] == '0' || digits[0] == '1' {
		return res
	}
	dict := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

	f := dict[digits[0]-'2']
	for i := 0; i < len(f); i++ {
		res = append(res, string(f[i]))
	}
	for i := 1; i < len(digits); i++ {
		if digits[i] == '1' {
			return []string{}
		}
		temp := make([]string, 0)
		for j := 0; j < len(res); j++ {
			for _, v := range dict[digits[i]-'2'] {
				temp = append(temp, res[j]+string(v))
			}
		}
		res = temp
	}
	return res
}

func Solution2(digits string) []string {
	res := make([]string, 0)
	if len(digits) < 1 || digits[0] == '0' || digits[0] == '1' {
		return res
	}
	dict := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

	f := dict[digits[0]-'2']
	for i := 0; i < len(f); i++ {
		res = append(res, string(f[i]))
	}
	for i := 1; i < len(digits); i++ {
		if digits[i] == '1' {
			return []string{}
		}
		l := len(res)
		for j := 0; j < l; j++ {
			for _, v := range dict[digits[i]-'2'] {
				res = append(res, res[j]+string(v))
			}
		}
		res = res[l:]
	}
	return res
}

func Solution(digits string) []string {
	res := make([]string, 0)
	if len(digits) < 1 || digits[0] == '0' || digits[0] == '1' {
		return res
	}
	dict := map[byte][]byte{'2': {'a', 'b', 'c'}, '3': {'d', 'e', 'f'}, '4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'}, '6': {'m', 'n', 'o'}, '7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'}, '9': {'w', 'x', 'y', 'z'}}

	for i := 0; i < len(dict[digits[0]]); i++ {
		res = append(res, string(dict[digits[0]][i]))
	}
	for i := 1; i < len(digits); i++ {
		if digits[i] == '1' {
			return []string{}
		}
		l := len(res)
		for j := 0; j < l; j++ {
			for _, v := range dict[digits[i]] {
				res = append(res, res[j]+string(v))
			}
		}
		res = res[l:]
	}
	return res
}
