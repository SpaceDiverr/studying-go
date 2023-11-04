package gopractising

import (
	// "fmt"
	// "reflect"
	// "strconv"
)

// func main() {
// 	// 1 4 9 16 25 36 49 64 81 100 => 4
// 	// 1 4 9 16 25 36 49 64 81 100 121 144, 4 => 5
// 	fmt.Println(CountNumberInSlice(8, 6))  // 3
// 	fmt.Println(CountNumberInSlice(10, 1)) // 4
// 	fmt.Println(CountNumberInSlice(12, 4)) // 5 
// 	fmt.Println("заебись блять.")
// 	fmt.Println(getCumOnFace(8, 6))
// 	fmt.Println(getCumOnFace(10, 1))
// 	fmt.Println(getCumOnFace(12, 4))
// 	fmt.Println("заебись блять.")
// }

func CountNumberInSlice(n, a int) int {
	var count int

	for i := 1; i <= n; i++ {
		val := i * i

		for val > 0 {
			if val % 10 == a {
				count++
			}

			val /= 10
		}
	}

	return count
}

func getCumOnFace(n, a int) int {
	var count, number int
	for i := 1; i <= n; i++ {
		number = i * i
		for number > 0 {
			if number%10 == a {
				count++
			}
			number /= 10
		}
	}
	return count
}
