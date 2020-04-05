package restore_ip_addresses

var ip string
var res []string

func RestoreIPAddresses(s string) []string {
	res = make([]string, 0)
	if len(s) < 4 || len(s) > 12 {
		return res
	}
	ip = s
	backtrack(0, 4, "")
	return res
}

func backtrack(idx, count int, sb string) {
	if len(ip[idx:]) < count || len(ip[idx:]) > 3*count {
		return
	}
	count--
	for i := 1; i < 4; i++ {
		if len(ip[idx+i:]) < count {
			return
		}
		if count > 0 {
			if i > 1 && ip[idx] == '0' {
				continue
			}
			if i == 3 {
				if ip[idx] > '2' || (ip[idx] == '2' && ip[idx+1] > '5') ||
					(ip[idx] == '2' && ip[idx+1] == '5' && ip[idx+2] > '5') {
					return
				}
			}
			str := ip[idx : idx+i]
			for str[0] == '0' && len(str) > 1 {
				str = str[1:]
			}
			backtrack(idx+i, count, sb+str+".")
		} else {
			for len(ip[idx:]) > 1 && ip[idx] == '0' {
				return
			}
			if len(ip[idx:]) == 3 {
				if ip[idx] > '2' || (ip[idx] == '2' && ip[idx+1] > '5') ||
					(ip[idx] == '2' && ip[idx+1] == '5' && ip[idx+2] > '5') {
					continue
				}
			}
			res = append(res, sb+ip[idx:])
			break
		}
	}
}
