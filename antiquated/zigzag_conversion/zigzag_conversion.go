package zigzag_conversion

func Convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	divVal := numRows - 1
	res, idx := make([]byte, len(s)), 0
	for i := 0; i < numRows; i++ {
		for p := i; p < len(s); {
			res[idx] = s[p]
			idx++
			if p%divVal != 0 && (divVal-p%divVal)*2+p < len(s) {
				res[idx] = s[(divVal-p%divVal)*2+p]
				idx++
			}
			p += divVal * 2
		}
	}
	return string(res)
}
