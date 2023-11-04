package gopractising

import (
	// "fmt"
	"reflect"
	"sort"
	// "strings"
	// "time"
)


	// var k []int = []int{1, 2, 3}
	// var a1 []int = nil
	// a2 := []int{11*11, 121*121, 144*144, 19*19, 161*161, 19*19, 144*144, 19*19}
	// fmt.Println(CompDima([]int{121, 144, 19, 161, 19, 144, 19, 11}, []int{121, 14641, 20736, 361, 25921, 361, 20736, 361}))
	// fmt.Println(CompDanya([]int{121, 144, 19, 161, 19, 144, 19, 11}, []int{121, 14641, 20736, 361, 25921, 361, 20736, 361}))
	// fmt.Println(CompDanya(a1, a2))
	// i := k
	// e := k[:]
	// fmt.Println(k)
	// fmt.Println(i)
	// fmt.Println(e)
	// var s string = "abddc_c"
	// fmt.Println(strings.Title(s))


func CompDanya(arr1 []int, arr2 []int) bool {
	if arr1 == nil || arr2 == nil {
		return false
	}
	sort.Ints(arr1)
	sort.Ints(arr2)
	for i, v := range arr1 {
		arr1[i] = v * v
	}
	return reflect.DeepEqual(arr1, arr2)
}

func CompDima(array1 []int, array2 []int) bool {
	for i := 0; i < len(array1); i++ {
		for j := 0; j < len(array2); j++ {
			if array2[j] == array1[i]*array1[i] {
				break
			}

			if j == (len(array2)-1) && array2[j] != array1[i]*array1[i] {
				return false
			}
		}
	}

	return true
}

// Given two arrays a and b write a function comp(a, b) (orcompSame(a, b)) that checks whether the two arrays have the "same" elements, with the same multiplicities
// (the multiplicity of a member is the number of times it appears). "Same" means, here, that the elements in b are the elements in a squared, regardless of the order.

// Examples
// Valid arrays
// a = [121, 144, 19, 161, 19, 144, 19, 11]
// b = [121, 14641, 20736, 361, 25921, 361, 20736, 361]
// comp(a, b) returns true because in b 121 is the square of 11, 14641 is the square of 121, 20736 the square of 144, 361 the square of 19,
// 25921 the square of 161, and so on. It gets obvious if we write b's elements in terms of squares:

// a = [121, 144, 19, 161, 19, 144, 19, 11]
// b = [11*11, 121*121, 144*144, 19*19, 161*161, 19*19, 144*144, 19*19]
// Invalid arrays
// If, for example, we change the first number to something else, comp is not returning true anymore:

// a = [121, 144, 19, 161, 19, 144, 19, 11]
// b = [132, 14641, 20736, 361, 25921, 361, 20736, 361]
