package gopractising

import (
	"fmt"
	"math"
	"sort"
)

/*
Given a positive number n > 1 find the prime factor decomposition of n. The result will be a string with the following form :

 "(p1**n1)(p2**n2)...(pk**nk)"
with the p(i) in increasing order and n(i) empty if n(i) is 1.

Example: n = 86240 should return "(2**5)(5)(7**2)(11)"
*/

func isPrime(n int) bool { 
	for i := 2; i < n; i++ {  
		if n % i == 0 {
			return false
		}
	}
	return true
}

func PrimeFactorsdima(n int) string {
	var (
		divisors []int
		decomposition string
	)

	for n > 1 {
		for i := 2; i <= n; i++ {
			if n % i == 0 && isPrime(i) { 
				divisors = append(divisors, i)
				n /= i
			}
		}
	}

	for i := 0; i < len(divisors); i++ {
		if i > 0 && divisors[i] <= divisors[i - 1] {
			break
		}

		power := 0

		for j := 0; j < len(divisors); j++ { 
			if divisors[i] == divisors[j] {
				power++
			}
		}

		if power > 1 {
			decomposition += fmt.Sprintf("(%d**%d)", divisors[i], power)
		} else {
			decomposition += fmt.Sprintf("(%d)", divisors[i])
		}
	}

	return decomposition
}

func PrimeFactors(n int) string {
	var primePotentialFactors []int = Erato(n)
	var dict map[int]int = map[int]int{}
	var out1 string 
	

	for _, v := range primePotentialFactors {
		for n % v == 0 {
			_, ok := dict[v]
			if !ok {
				dict[v] = 1
			} else {
				dict[v] += 1
			}
			n /= v
			
		}
	}

	var buf []int = make([]int, 0, len(dict))

	for key := range dict {
		buf = append(buf, key)
	}
	sort.Ints(buf)

	for _, key := range buf {
		if dict[key] == 1 {
			out1 += fmt.Sprintf("(%v)", key)
		} else {
			out1 += fmt.Sprintf("(%v**%v)", key, dict[key])
		}
	}
	return out1
}

func RemoveEach(arr []int, n int) []int {

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == n {
			arr = RemoveAt(arr, i)
		}
	}

	return arr
}

func RemoveAt(arr []int, i int) []int {
	return append(arr[:i], arr[i+1:]...)
}

func CreateSequence(start int, end int) []int {

	var capAndLen int = end - start
	var outputSlice []int = make([]int, capAndLen)

	for i := 0; i < len(outputSlice); i++ {
		outputSlice[i] = i + start
	}

	return outputSlice
}


func Erato(n int) []int {

	var arr []int = CreateSequence(0, n)
	var emptied int = 1
	var nsqrt int = int(math.Ceil(math.Sqrt(float64(n))))

	for i := 2; i <= nsqrt; i++  {
		if arr[i] != 0 {
			for j := 2*i ; j < len(arr); j += i {
				if arr[j] != 0 {
					arr[j] = 0
					emptied++
				}
			}
		}
	}

	RemoveEach(arr, 0)

	return arr[1:len(arr) - emptied]
}

func EratoSieveAndPrimeFactorsOf(n int) ([]int, map[int]int) {

	var arr []int = CreateSequence(0, n)
	var emptied int = 1
	var nsqrt int = int(math.Ceil(math.Sqrt(float64(n))))
	var dict map[int]int

	for i := 2; i <= nsqrt; i++ {
		if arr[i] != 0 {
			if (n % arr[i] == 0) { // рассмотреть случай c другим делителем = n/i
				_, ok := dict[arr[i]]
				if !ok {
					dict[arr[i]] = 1
				}
				dict[arr[i]] += 1
			}
			for j := 2 * i; j < len(arr); j += i {
				if arr[j] != 0 {
					arr[j] = 0
					emptied++
				}
			}
		}
	}
	RemoveEach(arr, 0)

	return arr[1:len(arr) - emptied], dict
}
func SliceContains(arr []int, n int) bool { // for not sorted arrs
	for _, v := range arr {
		if v == n {
			return true
		} 
	}
	return false
}

func dictIntIntContains(dict map[int]int, n int ) bool {
	_, ok := dict[n]
	if ok {
		return true
	} else {
		return false
	}
}
func dictStringIntContains(dict map[string]int, n string ) bool {
	_, ok := dict[n]
	if ok {
		return true
	} else {
		return false
	}
}
func dictStringStringContains(dict map[string]string, n string ) bool {
	_, ok := dict[n]
	if ok {
		return true
	} else {
		return false
	}
}
func dictIntStringContains(dict map[int]string, n int ) bool {
	_, ok := dict[n]
	if ok {
		return true
	} else {
		return false
	}
}

func cumUnity(n int) string {
	if (isPrime(n)) {
		return fmt.Sprintf("(%v)", n)
	}
	var divs []int
	var end int = int(math.Ceil(math.Pow(float64(n), 0.5)))
	var out1 string
	var dict map[int]int = map[int]int{}
	for i := 2; i <= end; i++ {
		if isPrime(i) && n % i == 0 { 
			for n % i == 0 {
				temp := n/i
				if dictIntIntContains(dict, i) {
					dict[i]++
					if temp != 1 {
						if isPrime(temp) {
							if dictIntIntContains(dict, temp) {
								dict[temp]++
							} else {
								dict[temp] = 1
							}
						}
					}
				} else {
					dict[i] = 1
					if isPrime(temp) {
						dict[temp] = 1
					}
				}
				n /= i
			}
		}
	}
	for key := range dict {
		divs = append(divs, key)
	}
	sort.Ints(divs)
	for _, key := range divs {
		if dict[key] == 0 {
			continue
		} else if dict[key] == 1 {
			out1 += fmt.Sprintf("(%v)", key)
		} else {
			out1 += fmt.Sprintf("(%v**%v)", key, dict[key])
		}
	}
	return out1
}

func cumOn(n int) string {
	var (
		nCopy = n
		out1, out2 string
		pow1, pow2 int
		end int = int(math.Ceil(math.Pow(float64(n), 0.5)))

	)
	for i := 2; i <= end; i++ {
		pow1, pow2 = 0, 0
		if isPrime(i) && n % i == 0 {
			temp := nCopy / i
			ifPrime := isPrime(temp)
			for nCopy % i == 0 {
				if ifPrime && temp != i {
					pow2++
				}
				pow1++
				nCopy /= i 
			}
			if pow2 == 0 {
				continue
			} else if pow2 == 1 {
				out2 = fmt.Sprintf("(%v)", temp) + out2
			} else {
				out2 = fmt.Sprintf("(%v**%v)", temp, pow2) + out2
			}
			if pow1 == 0 {
				continue
			} else if pow1 == 1 {
				out1 += fmt.Sprintf("(%v)", i)
			} else {
				out1 += fmt.Sprintf("(%v**%v)", i, pow1)
			}
		}
	}
	return out1 + out2
}

