package excel_sheet_column_title

func ConvertToTitle(n int) string {
	res := make([]byte, 0)
	for n > 0 {
		v := n % 26
		if v == 0 {
			v = 26
			n -= 1
		}
		res = append(res, byte(v-1)+'A')
		n /= 26
	}
	p1, p2 := 0, len(res)-1
	for p1 < p2 {
		res[p1], res[p2] = res[p2], res[p1]
		p1++
		p2--
	}
	return string(res)
}
