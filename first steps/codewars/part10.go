package gopractising

// Digital root is the recursive sum of all the digits in a number.
// Given n, take the sum of the digits of n. If that value has more than one digit, continue reducing in this way until a single-digit number is produced. The input will be a non-negative integer.

// Examples
//     16  -->  1 + 6 = 7
//    942  -->  9 + 4 + 2 = 15  -->  1 + 5 = 6
// 132189  -->  1 + 3 + 2 + 1 + 8 + 9 = 24  -->  2 + 4 = 6 //
// 493193  -->  4 + 9 + 3 + 1 + 9 + 3 = 29  -->  2 + 9 = 11  -->  1 + 1 = 2

func ReverseSlice(arr []int) []int {
	for i, j := 0, len(arr) - 1; i < j; i, j = i + 1, j - 1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func numberToSliceOfDigits(n int) []int {
	var out []int 
	for n > 0 {
		out = append(out, n%10)
		n /= 10
	}
	return ReverseSlice(out)
}


func getRecRoot(n int) int {
	var sum int = 0
	if n >= 10 {
		var arr []int = numberToSliceOfDigits(n)
		for _, v := range arr { sum += int(v) }
		if sum < 10 {
			return sum
		} else { return getRecRoot(sum)}
	} else { return n }
}

func DigitalRoot(n int) int { 	
	if n < 10 {
		return n
	} else {
		return DigitalRoot(CountDigits(n))
	}
} 

func CountDigits(n int) int {
	var sum int 
	
	for n > 0 {
		sum += n % 10
		n /= 10
	}

	return sum
}