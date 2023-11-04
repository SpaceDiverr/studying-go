package gopractising

import (
	"fmt"
	"sort"
)

type ByRune []rune

func (r ByRune) Len() int           { return len(r) }
func (r ByRune) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ByRune) Less(i, j int) bool { return r[i] < r[j] }

func StringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
			r = append(r, runeValue)
	}
	return r
}

func SortStringByCharacter(s string) string {
	var r ByRune = StringToRuneSlice(s)
	sort.Sort(r)
	return string(r)
}


func TwoToOne(s1 string, s2 string) (res string) {

	var s3 = []rune(s1 + s2)
	
	var runeDict = make(map[rune]int)

	for _, r := range s3 {
		if _, ok := runeDict[r]; !ok {
			runeDict[r] = 1
		}
	}

	s3 = make([]rune, 0)

	for r, _ := range runeDict {
		s3 = append(s3, r)
	}

	res = SortStringByCharacter(string(s3))

	fmt.Printf("res: %v\n", res)

	return res
}
