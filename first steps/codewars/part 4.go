package gopractising

import (

)

func stringReverse(s string) string {
	var pal string
	for _, v := range s {
		pal = string(v) + pal
	}
	return s
}