package multiply_strings

func MultiplyStrings(num1, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	sum := make([]uint8, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		var adv uint8
		for j := len(num2) - 1; j >= 0; j-- {
			mul := (num1[i]-'0')*(num2[j]-'0') + adv
			adv = mul / 10
			sum[i+j+1] += mul % 10
			if sum[i+j+1] > 9 {
				sum[i+j+1] = sum[i+j+1] % 10
				adv++
			}
		}
		if adv > 0 {
			sum[i] += adv
		}
	}
	if sum[0] == 0 {
		sum = sum[1:]
	}
	for i := 0; i < len(sum); i++ {
		sum[i] += '0'
	}
	return string(sum)
}
