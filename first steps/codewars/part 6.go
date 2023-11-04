package gopractising

// import "fmt"

// func main() {
// 	var dict map[int]int = map[int]int{}
// 	_, ok := dict[2]
// 	if dict[2] == 0 {
// 		fmt.Println(nil)
// 	}
// 	if !ok {
// 		dict[2] = 1
// 	} else {
// 		dict[2] += 1
// 	}
// 	fmt.Println(dict[2])
// }

func getSumOfConsecutiveOdds(NthRow int) int {
	return NthRow * NthRow * NthRow
}