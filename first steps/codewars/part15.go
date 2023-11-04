package gopractising

import (
	"math"
)

/*
1, 246, 2, 123, 3, 82, 6, 41 are the divisors of number 246. Squaring these divisors we get: 1, 60516, 4, 15129, 9, 6724, 36, 1681. The sum of these squares is 84100 which is 290 * 290.

Task
Find all integers between m and n (m and n integers with 1 <= m <= n) such that the sum of their squared divisors is itself a square.

We will return an array of subarrays or of tuples (in C an array of Pair) or a string.
The subarrays (or tuples or Pairs) will have two elements: first the number the squared divisors of which is a square and then the sum of the squared divisors.

Example:
list_squared(1, 250) --> [[1, 1], [42, 2500], [246, 84100]]
list_squared(42, 250) --> [[42, 2500], [246, 84100]]
The form of the examples may change according to the language, see "Sample Tests".

Note
In Fortran - as in any other language - the returned string is not permitted to contain any redundant trailing whitespace: you can use dynamically allocated character strings.

Link: https://www.codewars.com/kata/55aa075506463dac6600010d/train/go /cool =)
*/

func getDivisors(n int) []int {
	var divs []int 
	if n == 1 {return []int{1, 1}}
	for i := 1; i < int(math.Ceil(math.Sqrt(float64(n)))); i++ {
		if n % i == 0 {
			divs = append(divs, int(math.Pow(float64(i), 2))); 
			divs = append(divs, int(math.Pow(float64(n/i), 2)))
		}
	}
	return divs
}

func ListSquaredDanya(m, n int) [][]int{
	var divs [][]int = [][]int{}
	if n == 1 && m == 1 {return append(divs, []int{1, 1})}
	for i := m; i <= n; i++ {
		val := 0.0
		for j := 1; j <= int(math.Sqrt(float64(i))); j++ {
			if i % j == 0 && j * j != i || j == 1 && i == 1 {
				val += math.Pow(float64(j), 2) + math.Pow(float64(i/j), 2)
			}
		}
		sqrt := int(math.Sqrt(val))
		if sqrt * sqrt == int(val){
			divs = append(divs, []int{i, int(val)})
		} else if val == 2 {
			divs = append(divs, []int{1, 1})
		}
	}
	return divs
}


func ListSquaredAlpha(m, n int) [][]int {
	return [][]int{}

}

func ListSquared(m, n int) [][]int {
	ans := [][]int{}
	
	for i := m; i <= n; i++ {
		div := GetDivisors(i); 
		square_sum := 0

	  	for _, k := range div {
			square_sum += k * k
		}

		if IsSquare(square_sum) {
			ans = append(ans, []int{i, square_sum})
		}
	}

	return ans
}
  
func GetDivisors(n int) (div []int) { 
	i := 1
  
	for i * i <= n {
		if n % i == 0{
			div = append(div, i)
			if i != n / i {
				div = append(div, n/i)
			}
		}
		i += 1
	}
	
	return 
}
  
func IsSquare(n int) bool {
	if n % 3 == 2 {
		return false
	}
	
	div := GetDivisors(n)
	for _, k := range div {
		if k * k == n {
			return true
		}
	}
	return false
}