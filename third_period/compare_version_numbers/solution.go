package compare_version_numbers

func CompareVersion(version1, version2 string) int {
	p1, p2 := 0, 0
	for p1 < len(version1) || p2 < len(version2) {
		num1, num2 := 0, 0
		for p1 < len(version1) && version1[p1] != '.' {
			num1 = num1*10 + int(version1[p1]-'0')
			p1++
		}
		for p2 < len(version2) && version2[p2] != '.' {
			num2 = num2*10 + int(version2[p2]-'0')
			p2++
		}
		if num1 > num2 {
			return 1
		} else if num1 < num2 {
			return -1
		}
		p1++
		p2++
	}
	return 0
}
