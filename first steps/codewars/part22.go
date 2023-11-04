package gopractising

/*
Given an array of positive or negative integers

I= [i1,..,in]

you have to produce a sorted array P of the form

[ [p, sum of all ij of I for which p is a prime factor (p positive) of ij] ...]

P will be sorted by increasing order of the prime numbers. The final result has to be given as a string in Java, C#, C, C++ and as an array of arrays in other languages.

Example:
I = (/12, 15/); // result = "(2 12)(3 27)(5 15)"
[2, 3, 5] is the list of all prime factors of the elements of I, hence the result.

Notes:

It can happen that a sum is 0 if some numbers are negative!
Example: I = [15, 30, -45] 5 divides 15, 30 and (-45) so 5 appears in the result, the sum of the numbers for which 5 is a factor is 0 so we have [5, 0] in the result amongst others.

*/
import (
	"fmt"
	"math"
	"sort"
)

// heil totenest
// fmt.Println(-3 % 2)
// fmt.Println(GetPrimeFactors(200))
// fmt.Println(SumOfDividedCodeCummissar([]int{12, 15}))
// fmt.Println("must be equal to")
// fmt.Println("(2 12)(3 27)(5 15)")
// fmt.Println(SumOfDividedCodeCummissar([]int{15,21,24,30,45}))
// fmt.Println("must be equal to")
// fmt.Println("(2 54)(3 135)(5 90)(7 21)")
// fmt.Println(SumOfDivided([]int{12, 15}))
// fmt.Println("must be equal to")
// fmt.Println("(2 12)(3 27)(5 15)")
// fmt.Println(SumOfDivided([]int{15,21,24,30,45}))
// fmt.Println("must be equal to")
// fmt.Println("(2 54)(3 135)(5 90)(7 21)")

func IsPrime(n int) bool {
	s := int(math.Sqrt(float64(n)))
	for i := 0; i < s; i++ {
		if i%s == 0 {
			return false
		}
	}
	return true
}

func Max(lst []int) int {
	max := lst[0]
	for _, v := range lst {
		if max < v {
			max = v
		}
	}
	return max
}
func Min(lst []int) int {
	min := lst[0]
	for _, v := range lst {
		if min < v {
			min = v
		}
	}
	return min
}

func SumOfDividedCodeCummissar(lst []int) string {
	sort.Ints(lst)
	primes := Erato(lst[len(lst)-1])
	result := ""

	for _, p := range primes {
		var sum, count int
		for _, v := range lst {
			if v%p == 0 {
				count++
				sum += v
			}
		}
		if count > 0 {
			result += fmt.Sprintf("(%d %d)", p, sum)
		}
	}

	return result
}

func SumOfDivided(lst []int) (p string) {
	var pf []int

	for i := 0; i < len(lst); i++ {
		pf = append(pf, GetPrimeFactors(Abs(lst[i]))...)
	}

	for i := 0; i < len(pf); i++ {
		c := pf[i]

		for j := i; j < len(pf); j++ {
			if c == pf[j] || pf[j]%c == 0 {
				pf[j] = c
			}
		}
	}

	sort.Ints(pf)
	lastdiv := 0

	for i := 0; i < len(pf); i++ {
		sum := 0

		if pf[i] != lastdiv {
			lastdiv = pf[i]
		} else {
			continue
		}

		for j := 0; j < len(lst); j++ {
			if lst[j]%pf[i] == 0 {
				sum += lst[j]
			}
		}

		p = fmt.Sprintf("%s(%d %d)", p, pf[i], sum)
	}

	return p
}

func GetPrimeFactors(n int) []int {
	factors := []int{}
	i := 2

	for i*i <= n {
		if n%i == 0 {
			factors = append(factors, i)
			n /= i
		} else {
			i++
		}
	}

	if n > 1 {
		factors = append(factors, n)
	}

	return factors
}

func Abs(n int) int {
	if n >= 0 {
		return n
	} else {
		return -n
	}
}
