package gopractising

import (
	"fmt"
	"strconv"
	// "fmt"
)





func UintToBits(n uint) uint {
	var out string
	var binary int

	for n > 0 {
		out = fmt.Sprint(n % 2) + out
		n /= 2
	}

	binary, _ = strconv.Atoi(out)

	return uint(binary)
}

func CountBits(n uint) int {
	var count int

	for n > 0 {
		if n % 2 == 1 {
			count++
		}

		n /= 2
	}

	return count
}
