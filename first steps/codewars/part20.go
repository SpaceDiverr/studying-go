package gopractising

import (
	"fmt"
)

/* A format for expressing an ordered list of integers is to use a comma separated list of either
	* individual integers
	* or a range of integers denoted by the starting integer separated from the end integer in the range by a dash, '-'.
	The range includes all integers in the interval including both endpoints. It is not considered a range unless it spans at least 3 numbers. For example "12,13,15-17"
Complete the solution so that it takes a list of integers in increasing order and returns a correctly formatted string in the range format.
Example:
solution([-10, -9, -8, -6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20]);
// returns "-10--8,-6,-3-1,3-5,7-11,14,15,17-20"
*/


	// fmt.Println([]int{-6,-3,-2,-1,0,1,3,4,5,7,8,9,10,11,14,15,17,18,19,20}, "\n")
	// fmt.Println(SolutionAlphaPatternsIncluded([]int{-6,-3,-2,-1,0,1,3,4,5,7,8,9,10,11,14,15,17,18,19,20}), "\n")
	// //fmt.Println(SolutionCummachine([]int{-6,-3,-2,-1,0,1,3,4,5,7,8,9,10,11,14,15,17,18,19,20}) == "-6,-3-1,3-5,7-11,14,15,17-20", "\n")

	// fmt.Println("must be equal to:\n")
	// fmt.Println("-6,-3-1,3-5,7-11,14,15,17-20")
	

func SolutionCummachine(list []int) string {
	s := ""
	count := 0
	
	for i := 1; i < len(list); i++ {

		if list[i] - list[i - 1] == 1 {
			if count == 0 {
				s += fmt.Sprintf("%v", list[i-1])
			}
			count += 1

		} else if list[i] - list[i - 1] > 1 {
			if count == 0 {
				s += fmt.Sprintf("%v,", list[i-1])
				
			} else if count == 1 {
				s += fmt.Sprintf(",%v,", list[i-1])

			} else {
				s += fmt.Sprintf("-%v,", list[i-1])
			}
			count = 0
		}

		if i == len(list) - 1 {
			if len(list) > 1 && list[i] - list[i - 1] == 1 && list[i - 1] - list[i - 2] == 1 {
				s += fmt.Sprintf("-%v", list[i])

			} else if i == len(list) - 1 && list[i] - list[i - 1] >= 1 {
				if count == 0 {
					s += fmt.Sprintf("%v", list[i])
				} else {
					s += fmt.Sprintf(",%v", list[i])
				}
			}
			count = 0
		}
	}
	return s
}

func SolutionAlphaPatternsIncluded(list []int) (sol string) {
	//var test []int
	var count int
	fl := -1
	for i := 0; i < len(list) - 1; i++ {
		
		if (list[i + 1] - list[i] == 1) && (fl == -1) {
			fl = i
			//  fmt.Println(i, count)
			count = 1
			//count++
		} else if (list[i + 1] - list[i] == 1) && (fl != -1) {
			count++
		} else if (list[i + 1] - list[i] == 1) && (i == len(list) - 2) {
			sol = fmt.Sprintf("%s%d-%d", sol, list[fl], list[len(list)]);
		} else if (list[i + 1] - list[i] != 1) && (count > 1) {
			sol = fmt.Sprintf("%s%d-%d", sol, list[fl], list[i]);
			if i != len(list) - 1 {
				sol += ","
			}
			fl = -1
			count = 0
		} else if ((list[i + 1] - list[i] != 1) || (i == len(list) - 1)) && (count == 1) {
			sol = fmt.Sprintf("%s%d,%d", sol, list[fl], list[i]);
			if i != len(list) - 1 {
				sol += ","
			}
			fl = -1
			count = 0 
		} else {
			count = 0
			sol = fmt.Sprintf("%s%d", sol, list[i]);
			if i != len(list) - 1 {
				sol += ","
			}
		}
		fmt.Println(i, list[i], count)
		//test = append(test, count) 
	}
	//fmt.Println(test)
	return 
}