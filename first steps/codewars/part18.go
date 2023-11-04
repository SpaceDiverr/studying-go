package gopractising

import (
	// "fmt"
	"reflect"
	"strings"
)

// b1 := []string{"NORTH", "WEST", "EAST", "SOUTH", "EAST", "WEST"}
// fmt.Println(DirReduc(b1))

func CheckIfOpposite(a, b string) bool {
	var WE, NS = [2]string{"west", "east"}, [2]string{"north", "south"}
	a, b = strings.ToLower(a), strings.ToLower(b)
	if a == WE[0] && b == WE[1] ||
		a == WE[1] && b == WE[0] ||
		a == NS[0] && b == NS[1] ||
		a == NS[1] && b == NS[0] {
		return true
	}
	return false
}

func delRangeOfStrings(arr []string, start, end int) []string {
	if start > end {
		return []string{}
	}
	return append(arr[:start], arr[end+1:]...)
}

// func DirReducUNDONE(arr []string) []string {
// 	var buf []string
// 	var recFlag = true
// 	var i int
// 	for i < len(arr) - 1 {
// 		if CheckIfOpposite(arr[i], arr[i + 1]) {recFlag = false; i += 2} else {buf = append(buf, arr[i]); i++}
// 	}
// 	if !CheckIfOpposite(arr[len(arr)-1], arr[len(arr)-2]) && len(arr) >= 2 {buf = append(buf, arr[len(arr) - 1])}
// 	if recFlag {return arr} else {return DirReduc(buf)}
// }

func DirReduc(arr []string) []string {
	var recFlag = true
	var buf []string
	if reflect.DeepEqual(arr, []string{}) {
		return arr
	}
	for i := 0; i < len(arr)-1; i++ {
		if CheckIfOpposite(arr[i], arr[i+1]) {
			recFlag = false
			arr[i], arr[i+1] = "", ""
		}
	}
	if recFlag {
		return arr
	}
	for _, v := range arr {
		if v != "" {
			buf = append(buf, v)
		}
	}
	return DirReduc(buf)
}
