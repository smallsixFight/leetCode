package simplify_path

import "strings"

func SimplifyPath(path string) string {
	list := make([]string, 0)
	res := strings.Builder{}

	for i := 0; i < len(path); i++ {
		for path[i] == '/' && i+1 < len(path) && path[i+1] == '/' {
			i++
		}
		p := make([]byte, 0)
		if path[i] != '/' {
			p = append(p, '/')
		} else {
			p = append(p, path[i])
			i++
		}
		pointNum := 0
		for i < len(path) {
			if path[i] == '.' {
				pointNum++
			} else if path[i] == '/' {
				break
			}
			p = append(p, path[i])
			i++
		}
		if pointNum == 2 && len(p) == 3 {
			if len(list) > 0 && list[len(list)-1] != "/" {
				list = list[:len(list)-1]
			}
		} else if pointNum == 1 {
			if len(p) > 2 {
				list = append(list, string(p))
			}
		} else if len(p) > 1 {
			list = append(list, string(p))
		}
	}
	for k := range list {
		res.WriteString(list[k])
	}
	r := res.String()
	if r == "" {
		return "/"
	}
	return r
}
