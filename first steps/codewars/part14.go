package gopractising

/*
You are given an array (which will have a length of at least 3, but could be very large) containing integers.
The array is either entirely comprised of odd integers or entirely comprised of even integers except for a single integer N.
Write a method that takes the array as an argument and returns this "outlier" N.

Examples
[2, 4, 0, 100, 4, 11, 2602, 36]
Should return: 11 (the only odd number)

[160, 3, 1719, 19, 11, 13, -21]
Should return: 160 (the only even number)
*/

// ссылка на задачу:
// https://www.codewars.com/kata/5526fc09a1bbd946250002dc/train/go


/* это тесты 
	Expect(FindOutlier([]int{2, 6, 8, -10, 3})).To(Equal(3))
	Expect(FindOutlier([]int{206847684, 1056521, 7, 17, 1901, 21104421, 7, 1, 35521, 1, 7781})).To(Equal(206847684))
	Expect(FindOutlier([]int{math.MaxInt32, 0, 1})).To(Equal(0))
*/ 
func FindOutlierDanya(ints []int) int { // это причем kyu 1))))))))
	var (
		flgOdd, flgOddFinal, flgEven, flgEvenFinal bool = false, false, false, false

	)
	for i, v := range ints {
		if flgOddFinal && flgEvenFinal {
			return v
		}
		if flgEven && v % 2 == 0 && !flgOdd {
			for j := i + 1; j < len(ints); j++ {
				if ints[j] % 2 == 1 {return ints[j]}
			}
		} else if flgOdd && v % 2 == 1 && !flgEven {
			for j := i + 1; j < len(ints); j++ {
				if ints[j] % 2 == 0 {return ints[j]}
			}
		}
		if v % 2 == 0 {flgEven = true} else {flgOdd = true}

	}
	return 0
}
// блять у этой хуйни даже авторское решение ебанутое боже даня на что мы подписались
func FindOutlierDanya2(ints []int) int { 
	for i := 0; i < len(ints) - 1; i++ {
		var cur, next = Abs(ints[i] % 2) == 0, Abs(ints[i + 1] % 2) == 0
		if (cur == next) {
			continue
		} else {
			if i == 0 {
				if cur == (ints[2] % 2 == 0){
					return ints[1]
				} else {
					return ints[0]
				}
			} else {return ints[i + 1]}
		}
	}
	return 0
}

/*
ну и че тут
напрограммированн
о 
*/ 
func FindOutlierDima(a []int) int {
	c := 0; d := 0; o := 0
	for i := 0; i < len(a); i++ {
		if a[i] % 2 == 0 {
			c++
			o = i
		} else {
			d = i
		}
	}
	if c == 1 {
		return a[o]
	} else {
		return a[d]
	}
}