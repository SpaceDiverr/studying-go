package gopractising

import (
	"fmt"
)

func CreateSequenceFromTo(start int, end int) []int {

	if start == end {
		return []int{start}
	}

	capAndLen := Abs(end) - Abs(start) + 1
	outputSlice := make([]int, capAndLen)

	if start > end {
		start = end
	}

	for i := 0; i < len(outputSlice); i++ {
		outputSlice[i] = i + start
	}

	return outputSlice
}

func SumOfDivided(lst []int) string {
	min, max := Abs(Min(lst)), Abs(Max(lst))

	if min > max {
		max = min
	}
	primes := Erato(max)
	s := ""

	for _, p := range primes {
		sum, flag := 0, false
		for _, v := range lst {
			if v%p == 0 {
				sum += v
				flag = true
			}
		}
		if flag {
			s += fmt.Sprintf("(%v %v)", p, sum)
		}
	}
	return s
}
