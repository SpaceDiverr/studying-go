package gopractising

import (
	// "fmt"
	"strings"
)

// Count the number of Duplicates
// Write a function that will return the count of distinct case-insensitive alphabetic characters and numeric digits that occur more than once in the input string. The input string can be assumed to contain only alphabets (both uppercase and lowercase) and numeric digits.

// Example
// "abcde" -> 0 # no characters repeats more than once
// "aabbcde" -> 2 # 'a' and 'b'
// "aabBcde" -> 2 # 'a' occurs twice and 'b' twice (`b` and `B`)
// "indivisibility" -> 1 # 'i' occurs six times
// "Indivisibilities" -> 2 # 'i' occurs seven times and 's' occurs twice
// "aA11" -> 2 # 'a' and '1'
// "ABBA" -> 2 # 'A' and 'B' each occur twice


func duplicate_countDanya(s string) int {
	var count int
	var dict map[rune]int = map[rune]int{}
	s = strings.ToLower(s)

	for _, v := range s {
		dict[v]++
	}
	for _, v := range dict {
		if v > 1 {
			count++
		}
	}
	return count
}

func duplicate_count(s1 string) int {
	var n int

	s1 = strings.ToLower(s1)

	// A B B A

	for i := 0; i < len(s1); i++ {
		if s1[i] == s1[0] {
			n++
		}
	}

	return n
}
