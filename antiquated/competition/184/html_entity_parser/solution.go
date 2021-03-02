package html_entity_parser

import "strings"

func EntityParser(text string) string {
	sb := strings.Builder{}
	for i := 0; i < len(text); {
		if text[i] == '&' {
			if i+5 < len(text) && text[i:i+6] == "&quot;" {
				sb.WriteByte('"')
				i += 6
			} else if i+5 < len(text) && text[i:i+6] == "&apos;" {
				sb.WriteByte('\'')
				i += 6
			} else if i+4 < len(text) && text[i:i+5] == "&amp;" {
				sb.WriteByte('&')
				i += 5

			} else if i+3 < len(text) && text[i:i+4] == "&gt;" {
				sb.WriteByte('>')
				i += 4
			} else if i+3 < len(text) && text[i:i+4] == "&lt;" {
				sb.WriteByte('<')
				i += 4
			} else if i+6 < len(text) && text[i:i+7] == "&frasl;" {
				sb.WriteByte('/')
				i += 7
			} else {
				sb.WriteByte(text[i])
				i++
			}
		} else {
			sb.WriteByte(text[i])
			i++
		}
	}
	return sb.String()
}
