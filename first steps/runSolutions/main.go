package main

import (
	"fmt"
	"math"
	"strings"
	"sort"
	// p "Tasks/gopractising"
	"bufio"
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

func SortStringByRune(s string) string {
	var r ByRune = StringToRuneSlice(s)
	sort.Sort(r)


	return string(r)
}


func TwoToOne(s1 string, s2 string) (res string) {

	var builder strings.Builder
	builder.WriteString(s1)
	builder.WriteString(s2)
	
	var runeDict = make(map[rune]int)

	for _, r := range builder.String() {
		if _, ok := runeDict[r]; !ok {
			runeDict[r] = 1
		}
	}

	builder = strings.Builder{}

	for r, _ := range runeDict {
		builder.WriteRune(r)
	}

	res = SortStringByRune(builder.String())

	return res
}



func FindNextSquare(sq int64) int64 {

	s := float64(sq)
	rooted := math.Sqrt(s)

	if (float64(int(rooted)) == s / rooted) {
		return int64(math.Pow(rooted + 1, 2))
	}

	return -1
}

func PrinterError(s string) string {
	count := 0
	for _, r := range(s) {
		if (r > 'm') {

			count += 1
		}
	}
	return fmt.Sprintf("%d/%d", count, len(s))
}

func AbbrevName(name string) string {
	
	var split []string = strings.Split(name, " ")

	return strings.Join([]string{
		strings.ToUpper(string(split[0][0])), 
		".", 
		strings.ToUpper(string(split[1][0])),
		}, "")
}

type Fighter struct {
    Name            string
    Health          int
    DamagePerAttack int
}

func DeclareWinner(fighter1 Fighter, fighter2 Fighter, firstAttackerName string) string {

	var (
		flag = false
	)

	if fighter1.Name == firstAttackerName {
		flag = true
	}

	for {
		if fighter1.Health > 0 && fighter2.Health > 0 {
			if flag {
				fighter2.Health -= fighter1.DamagePerAttack
			} else {
				fighter1.Health -= fighter2.DamagePerAttack
			}
			flag = !flag
		} else {
			if fighter1.Health <= 0 {
				return fighter2.Name
			}
			return fighter1.Name
		}
	}
}


// func (*int) HasYears(seconds int) bool {
// 	if seconds - sInY < 0 { return false }
// 	return true 
// }

func Contains(seconds int, timeAmount int) bool {
	return seconds - timeAmount >= 0
}

type Duration struct {
	// Has
	Time int
}

// type Has interface {
// 	Seconds(time time)
// 	Minutes(time time)
// 	Hours(time time)
// 	Days(time time)
// 	Years(time Duration)
// }

// func (*int) Seconds()  {

// }

func stringReverse(s string) string { 
	var pal string = ""
	for _, v := range s {
		pal = string(v) + pal
	}
	return pal
}

func UpdateIfPlural(i int, s string) string {
	if i == 1 {
		return s 
	}
	return s + "s"
}

func FormatDuration(seconds int) (res string) {

	if seconds == 0 { return "now." }

	const (
		sInY = 3600 * 24 * 365
		sInD = 3600 * 24
		sInH = 3600
		sInM = 60
	)
	TIME := []int{sInY, sInD, sInH, sInM}

	durations := []int{0, 0 , 0, 0, seconds}
	
	count := 0

	for i := 0; i < len(TIME); i++ {
		if Contains(seconds, TIME[i]) {
			durations[i] = durations[len(durations) - 1] / TIME[i]
			durations[len(durations) - 1] -= TIME[i] * durations[i]
		}
	}

	timeNames := [...]string{"year", "day", "hour", "minute", "second"}
	
	for _, el := range durations {
		if el != 0 {
			count++
		}
	}

	if count >= 2 {
		for i, e := range durations {

			if e != 0 {
				if count == 1 {
					res += fmt.Sprintf("and %d %s", durations[i], UpdateIfPlural(durations[i], timeNames[i]))
					break
				}
				if count == 2 {
					res += fmt.Sprintf("%d %s ", durations[i], UpdateIfPlural(durations[i], timeNames[i]))
					count--
					continue
				}
				count--
				res += fmt.Sprintf("%d %s, ", durations[i], UpdateIfPlural(durations[i], timeNames[i]))
			}
		}
	} else {
		for i, e := range durations {
			if e != 0 {
				res = fmt.Sprintf("%d %s", durations[i], UpdateIfPlural(durations[i], timeNames[i]))
				break
			}
		}
	}
	res += "."
	return 
}



func main() {

	// fmt.Println(FormatDuration(3600 * 24 * 365 * 2 + 24 * 3600 + 3600 * 5 + 1))
	// var str string = stringReverse("asdf")

	// println(str)
	// fmt.Println(str)

	fmt.Printf("TwoToOne: %v\n", TwoToOne("xxxxyyyyabklmopq", "xyaabbbccccdefww"))
 
}






