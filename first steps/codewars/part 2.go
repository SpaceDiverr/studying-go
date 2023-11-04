//package kata

package gopractising

/*
Task
Given three integers a ,b ,c, return the largest number obtained after inserting the following operators and brackets: +, *, ()
In other words , try every combination of a,b,c with [*+()] , and return the Maximum Obtained
*/

import (
	// f "fmt"
	s "sort"
)

// func main() {
// 	f.Println(AlphaPattern(1, 3, 1))     // 3
// 	f.Println(ExpressionMatter(2, 2, 2)) // 8
// 	f.Println(ExpressionMatter(3, 0, 5)) // 15
// 	f.Println(ExpressionMatter(5, 1, 3)) // 20
// 	f.Println(ExpressionMatter(2, 1, 2)) // 6
// }

func sum(s []int) int {
	var sum int
	for _, val := range s {
		sum += val
	}
	return sum
}

func searchIf(arr []int, num int) bool {
	for _, el := range arr {
		if el == num {
			return true
		}
	}
	return false
}

func searchInd(arr []int, num int, leftOrRight bool) int {

	var count int
	var out int = -1

	leftOrRight = true //left

	if leftOrRight { //left

		for i := 0; i < count; i++ {
			if arr[i] == num {
				out = i
				break
			}
		}

	} else { // right

		for i := len(arr) - 1; i >= 0; i-- {
			if arr[i] == num {
				out = i
				break
			}
		}
	}

	return out
}

func searchIndOfAll(arr []int, num int) []int {
	var out []int
	for i, v := range arr {
		if v == num {
			out = append(out, i)
		}
	}
	return out
}

func searchCountOf(arr []int, num int) int {
	var count int = 0
	for _, v := range arr {
		if v == num {
			count++
		}
	}
	return count
}

func product(s []int) int {
	var product int = 1
	for _, val := range s {
		product *= val
	}
	return product
}

func ExpressionMatter(a, b, c int) int {
	var unreachable int = -1
	var outOneElement int = 1
	var zeroElCount = 0
	var outProductOfTwo = 1
	var antiZeros []int

	if a == 0 && b == 0 && c == 0 {
		return 0
	}

	var arr []int = []int{a, b, c}
	s.Ints(arr)

	for i, el := range arr {
		if el == 0 {
			zeroElCount++

		} else {
			outOneElement += el
			outProductOfTwo *= el
			antiZeros = append(antiZeros, arr[i])

		}
	}

	if zeroElCount == 1 {
		return outProductOfTwo

	} else if zeroElCount == 2 {
		return outOneElement

	}

	switch len(antiZeros) {

	case 3:
		if searchCountOf(antiZeros, 1) == 3 {
			return sum(antiZeros)

		} else if searchCountOf(antiZeros, 1) >= 1 {
			return (antiZeros[0] + antiZeros[1]) * antiZeros[2]

		}
		return a * b * c

	case 2:
		return product(antiZeros)

	case 1:
		return sum(antiZeros)

	}
	return unreachable
}

func e(a, b int) bool {
	return a == b
}

func cum(a, b, c int) int {

	var (
		arr = []int{a, b, c}

		zeroCount = searchCountOf(arr, 0)
		oneCount  = searchCountOf(arr, 1)

		unreachableEnd = 0
	)

	if zeroCount >= 2 {
		return sum(arr)

	} else if zeroCount == 1 {

		if a == 0 {
			if oneCount == 0 {
				return b * c
			} else if oneCount == 1 {
				return b + c
			}
		} else if b == 0 {

			return a + c

		} else if c == 0 {
			if oneCount == 0 {
				return a * b
			} else if oneCount == 1 {
				return a + b
			}
		}
	} else if zeroCount == 0 {

		if oneCount == 0 {
			return product(arr)

		} else if oneCount == 1 {

			if b == 1 {
				if a > c {
					return a * (b + c)
				} else {
					return c * (b + a)
				}
			} else if a == 1 {
				return c * (b + a)

			} else if c == 1 {
				return a * (b + c)
			}
		} else if oneCount == 2 {
			if b != 1 {
				return sum(arr)
			} else if a != 1 {
				return a * 2
			} else if c != 1 {
				return c * 2
			}
		} else if oneCount == 3 {
			return 3
		}
	}
	return unreachableEnd
}

func AlphaPattern(a, b, c int) int {
	var result int

	// 1 3 1
	// 2 * 3
	// min(4, min(2, 4)) * max(1, max(3, 1))
	if a == 1 || b == 1 || c == 1 {
		result = max(a+b+c, min((a+b), min(a+c, b+c))*max(a, max(b, c)))
	} else if a*b*c == 0 {
		result = max(a, b) * max(a, c)
	} else {
		result = a * b * c
	}

	return result
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
