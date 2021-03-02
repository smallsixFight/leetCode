package add_binary

func AddBinary(a string, b string) string {
	if len(a) == 0 {
		return a
	} else if len(b) == 0 {
		return b
	}
	l := 0
	if len(a) > len(b) {
		l = len(a)
	} else {
		l = len(b)
	}
	arr := make([]uint8, l+1)
	i, j := len(a)-1, len(b)-1
	var plusVal uint8
	for l >= 0 {
		var temp uint8
		if i >= 0 && j >= 0 {
			temp = a[i] + b[j] + plusVal
		} else if i >= 0 {
			temp = a[i] + plusVal
		} else if j >= 0 {
			temp = b[j] + plusVal
		} else {
			temp = '0' + plusVal
		}
		if temp%48 > 1 {
			plusVal = 1
		} else {
			plusVal = 0
		}
		arr[l] = temp%2 + '0'
		i--
		j--
		l--
	}
	if arr[0] == '0' {
		return string(arr[1:])
	}
	return string(arr)
}
