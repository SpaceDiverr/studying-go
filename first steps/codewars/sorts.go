package gopractising

import (
	// "fmt"
	// "math"
	"math/bits"
	"math/rand"
	// "time"
	"math"
	// "practisinggo/intMath"
)

// import "github.com/go-sql-driver/mysql"


	// algTime := time.Now()
	// sliceLength := math.Pow10(5)
	// fmt.Println(selectionSort(CreateSliceOfRandInts(uint(sliceLength), maxVal)), "\n")
	// timePassed := time.Since(algTime).Seconds()
	// fmt.Println("\"selectionSort\" function call took", timePassed, "seconds.")

	// var maxVal uint = 2000
	// fmt.Println("------------------------")
	// var intSlice *IntSlice = new(IntSlice)
	// intSlice.container = CreateSliceOfRandInts(50, maxVal)

	// fmt.Println(intSlice.CountFunc(IsIntEven))
	// fmt.Println(intSlice.container)
	// intSlice.container = []int{}
	// a := []int{1, 2, 3,}
	// b := make([]int, len(a))
	// c := copy(b, a)
	// fmt.Println(c == len(a))



func IsIntEven(n int) (res bool) {return n > 0 && n % 2 == 0}

type IntSlice struct {
	container []int 
}

type SliceFuncs interface {
	CountF(f func() bool)
	CountAll(n int)
}

func (slice *IntSlice) CountFunc(f func(n int) bool) (count int){
	for _, v := range slice.container {
		if f(v) {count++}
	}
	return
}

func (slice *IntSlice) CountAll(n int) (count int) {
	for _, v := range slice.container {
		if v & n == 0 {count++}
	}
	return
}

// uniform returns a uniformly random float in [0,1).
func uniform(precision int) int {
	if precision == 0 { precision = 1000}
	sig := rand.Uint64() % (1 << 52)
	return int(((1 + float64(sig)/(1<<52)) / math.Pow(2, geometric())) * float64(precision))
}
// geometric returns a number picked from a geometric
// distribution of parameter 0.5.
func geometric() float64 {
	b := 1.0
	for rand.Uint64()%2 == 0 {
		b++

	}
	return b
}

// get fully random sign (+ or -) for a given number based on how many bytes it takes on the average to represent "num" 
func GetRandomSign(num, bytes int) (res int) { // e.g. num := 1000
	// rangeLimit := bytes * 3 // num == 1000 -> 2 bytes -> rangeLimit == 6; rangeLimit is [0..6)
	bytesToMove := uniform(10) // here we take random number of bits to move in num at the next if statement
	res = num // binary(num) == 1111101000

	if (num >> bytesToMove) % 2 == 0 { 
		res = -num
	}
	return
}

func CreateSliceOfRandInts(length uint, maxVal uint) []int {
	out := make([]int, length)
	var curNum int
	// emptying used memory since func returns
	defer func() { curNum = 0; out = []int{} } ()

	minBytes := bits.Len(maxVal)

	for i := range out{
		curNum = uniform(1000)
		out[i] = GetRandomSign(curNum, minBytes) 
	}

	return out
}

func FindMinAndMinInd(arr []int) (min, minInd int) {
	min = arr[0]
	for i, v := range arr {
		if min > v {
			min = v
			minInd = i
		}
	}
	return 
}

func selectionSort(arr []int) (res []int) {
	for len(arr) > 0 {
		min, minInd := FindMinAndMinInd(arr)
		res = append(res, min)
		arr = append(arr[:minInd], arr[minInd + 1:]...)
	}
	return
}

func bubbleSort(arr []int) (res []int) {
	for i := 0; i < len(arr); i++ {
		min, minInd := FindMinAndMinInd(arr[i:])
		arr[i], arr[minInd] = min, arr[i]
	}
	return
}