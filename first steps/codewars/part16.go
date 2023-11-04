package gopractising

import (
	"fmt"
	"math"
	"strconv"
	"strings"


	// "regexp"
)
const (
	b1 = `
1000.00!=

125 Market !=:125.45
126 Hardware =34.95
127 Video! 7.45
128 Book :14.32
129 Gasoline ::16.10
`
	b1sol = "Original Balance: 1000.00\n125 Market 125.45 Balance 874.55\n126 Hardware 34.95 Balance 839.60\n127 Video 7.45 Balance 832.15\n128 Book 14.32 Balance 817.83\n129 Gasoline 16.10 Balance 801.73\nTotal expense  198.27\nAverage expense  39.65"

)

func delAllOf(s string, comp []rune) string {
	if ContainsAnyCharOf(s, comp) {
		for _, r := range comp {
			for strings.ContainsRune(s, r) {
				s = string(delCharAt([]rune(s), strings.Index(string(s), string(r))))
			}
		}
	} 
	return s
}

func recDelAllOf(s string, comp []rune ) string {
	if ContainsAnyCharOf(s, comp) { 
		for _, r := range comp {
			if strings.ContainsRune(string(s), r) {
				s = string(delCharAt([]rune(s), strings.Index(s, string(r))))
			}
		}
	}
	if ContainsAnyCharOf(s, comp) {return recDelAllOf(s, comp)}
	return s
}
func ContainsCharAndItsInd(s string, r rune) (bool, int) {
	for i, v := range s {
		if v == r {
			return true, i
		}
	}
	return false, -1
}
func ContainsAnyCharOf(root string, sub []rune) bool {
	for _, v := range sub {
		for _, r := range root {
			if r == v {
				return true
			}
		}
	}
	return false
}
func delCharAt(s []rune, index int) []rune {
	return append(s[0:index], s[index+1:]...)
}
func delStringAt(arr []string, i int) []string {
	return append(arr[0:i], arr[i+1:]...)
}
func BD(book string) string {
	sol := ""
	str := strings.FieldsFunc(book, func(r rune) bool { return r == ' ' || string(r) == "\n" || string(r) == "\t" })
	comp := []rune(":=;!?{},")
	var cash, avg, total float64
	var count int 

	fmt.Println(str)
	for i := range str {
		str[i] = delAllOf(str[i], comp)
	}
	fmt.Println(str)
	sol = fmt.Sprintf("Original Balance: %s\n", str[0])
	buf, _ := strconv.ParseFloat(str[0], 64)
	cash = math.Floor(buf * 100 / 100)
	str = delStringAt(str, 0)
	for i := 0; i <= len(str)-3; i += 3 {
		buf, _ := strconv.ParseFloat(str[i + 2], 64)
		count++
		avg += buf
		cash -= buf
		a, _ := strconv.ParseFloat(str[i+2], 64)
		

		sol += fmt.Sprintf("%s %s %s Balance %v\n", str[i], str[i+1], fmt.Sprintf("%.2f", a), fmt.Sprintf("%.2f", cash))
	}
	total = avg
	avg = avg / float64(count)
	sol += fmt.Sprintf("Total expense  %v\nAverage expense  %v", fmt.Sprintf("%.2f", total), fmt.Sprintf("%.2f", avg))
	fmt.Println(sol)
	return sol
}

func BalanceDanya(book string) []string {
	var buf []string
	// var i = 0
	for i := 0; i < len(book); i++ {
		var indCount int
		var j int = i
		buf = append(buf, book[i:i+indCount])
		if j+1 < len(book) {
			i = j + 1
		}
	}
	return buf
}

func BalanceDanyaRegxp(book string) { // string
	// modified := ``
	// // reLines := regexp.MustCompile(`\d+\s\w+\[!"#$%&'(),\-./:;<=>?@]^_{|}~]\s\d+\.\d+\[!"#$%&'(),\-./:;<=>?@[\\\]^_{|}~]*\n`)
	// reHeader := regexp.MustCompile(`\d{4}\.\0{2}`)
	// // lines := reLines.FindAllString(book, -1)
	// header := reHeader.FindString(book)

	// fmt.Println(lines)
	// fmt.Println(header)
	return
}

func BalanceDima(book string) string {
	// your code
	return ""
}
