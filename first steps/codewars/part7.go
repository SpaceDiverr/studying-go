package gopractising

import (
	
	"strings"
)


func Accum(s string) string {
	var buf string
	for i := 0; i < len(s); i++ {
		buf += strings.ToUpper(string(s[i]))
		for j := 0; j < i; j++ {
			buf += strings.ToLower(string(s[i]))
		}
		if (i != len(s) - 1) {
			buf += "-"
		}
	}
	return buf
}
