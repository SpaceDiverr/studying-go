package gopractising

/*
What is an anagram? Well, two words are anagrams of each other if they both contain the same letters. For example:
    'abba' & 'baab' == true
    'abba' & 'bbaa' == true
    'abba' & 'abbba' == false
    'abba' & 'abca' == false
Write a function that will find all the anagrams of a word from a list.
You will be given two inputs a word and an array with words. You should return an array of all the anagrams or an empty array if there are none. For example:
    anagrams('abba', ['aabb', 'abcd', 'bbaa', 'dada']) => ['aabb', 'bbaa']
    anagrams('racer', ['crazer', 'carer', 'racar', 'caers', 'racer']) => ['carer', 'racer']
    anagrams('laser', ['lazing', 'lazy',  'lacer']) => []
Note for Go
For Go: Empty string slice is expected when there are no anagrams found.
*/

import (
	// "fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)


// 	fmt.Println(Anagrams("carer", []string{"carer", "arcre", "carre", "racers", "arceer", "raccer", "carrer", "cerarr"}))
// 	fmt.Println("[" + "original" + `]{` + strconv.Itoa(len("original")) + `}`)
// 	fmt.Println(IsNotAnagram("abb", "bab"))
// 	fmt.Println(AnagramsDanya("abba", []string{"aabb", "abcd", "bbaa", "dada"}))
	// fmt.Println(strings.Split("abc", ""))
	// fmt.Println(len("regexp.MustCompile(`[codeITez]{8}`)"))


func IsNotAnagram(original string, toCumpare string) (res bool) {
	length := strconv.Itoa(len(original))
	re := regexp.MustCompile(`[` + original  + `]{` + length + `}`)
	return re.MatchString(toCumpare)
}

func AnagramsDanya(word string, words []string) []string {
	arr := []string{}

	for i, v := range words {
		if len(v) > len(word) {continue}
		if IsNotAnagram(word, v) {
			arr = append(arr, words[i])
		}
	}
	if arr == nil || len(arr) == 0 {return nil}
	return arr
}

func ContainsRune(arr string, s rune) bool {
	for _, r := range arr {
		if r == s {return true}
	}
	return false
}
func Anagrams1(word string, words []string) []string {
	arr := []string{}
	for _, str := range words {
		if len(str) > len(word) {continue}
		if reflect.DeepEqual(strings.Split(str, ""), strings.Split(word, "")) {arr = append(arr, str)}
	}
	if len(arr) == 0 || arr == nil {return nil}
	return arr
}

func Anagrams(word string, words []string) []string {
	arr := []string{}
	for _, v := range words {
    if len(v) > len(word) {
      continue
    }
		flag := true
		for _, r := range v {
			if !ContainsRune(word, r) {
				flag = false
			}
		}
		if flag {
			arr = append(arr, v)
		}
	}
	if len(arr) == 0 || arr == nil {return nil}
	return arr
  
}


func AnagramsCumsForever(word string, words []string) []string {
	arr := []string{}
	for i, v := range words {
		if reflect.DeepEqual(v, word){
			arr = append(arr, words[i])
		}
	}
	if reflect.DeepEqual(arr, []string{}) {return nil}
	return arr
}

func AnagramsDima(word string, words []string) (ans []string) {
    var k map[byte]int
    k = make(map[byte]int)
    
    for i := 0; i < len(word); i++ {
        k[word[i]]++
    }

    for i := 0; i < len(words); i++ {
        var c map[byte]int
        c = make(map[byte]int)
        
        for j := 0; j < len(words[i]); j++ {
            c[words[i][j]]++
        }
        
        for j := 0; j < len(word); j++ {
            if c[word[j]] != k[word[j]] || len(c) != len(k) {
                break
            }

            if j == len(word) - 1 {
                ans = append(ans, words[i])
            }
        }
    }
    
    return
}