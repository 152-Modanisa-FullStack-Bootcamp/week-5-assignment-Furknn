package assignment

import (
	"math"
	"sort"
	"strings"
)

func AddUint32(x, y uint32) (uint32, bool) {
	result := x + y
	overflow := false

	//If result is smaller than both inputs, it is oveflowed
	//Would not work for signed integers
	if result < x && result < y {
		overflow = true
	}
	return result, overflow
}

func CeilNumber(f float64) float64 {

	// Get mod of input for 0.25
	mod := math.Mod(f, 0.25)
	if mod != 0 {
		return f + (0.25 - mod)
	} else {

		//If mod is zero, return input without any changes
		return f
	}

}

func AlphabetSoup(s string) string {
	var charArr []rune

	//Convert string to rune array
	for _, val := range s {
		charArr = append(charArr, val)
	}

	//Sort
	sort.Slice(charArr, func(i, j int) bool {
		return charArr[i] < charArr[j]
	})

	//Convert array back to string and return
	return string(charArr)
}

func StringMask(s string, n uint) string {
	var charArr []rune

	if len(s) == 0 {
		return "*"
	}

	//Convert string to rune array
	for _, val := range s {
		charArr = append(charArr, val)
	}
	for i := range charArr {

		//n smaller than length of string, start from index
		if n < uint(i+1) {
			charArr[i] = '*'
		}

		//n is zero, bigger or equal to lengt of string, make all letters "*"
		if n == 0 || int(n) >= len(s) {
			charArr[i] = '*'
		}
	}
	return string(charArr)
}

func WordSplit(arr [2]string) string {
	foundWords := make(map[int]string)
	result := ""
	words := strings.Split(arr[1], ",")
	s := arr[0]

	//check string for words, if any found add to foundWords with its index as key
	for _, word := range words {
		if strings.Contains(s, word) {
			foundWords[strings.Index(s, word)] = word
		}
	}

	// if found less than two, return "not possible"
	if len(foundWords) < 2 {
		return "not possible"
	}

	// sort foundWords by key
	keys := make([]int, 0, len(foundWords))
	for k := range foundWords {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	//create response string
	for i, key := range keys {
		result = result + foundWords[key]

		//if not at last index, add comma
		if i+1 != len(keys) {
			result = result + ","
		}
	}
	return result
}

//Contains function for variadicSet
func variadicSetContains(slice []interface{}, value interface{}) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

func VariadicSet(i ...interface{}) []interface{} {
	var result []interface{}
	for _, val := range i {
		if !variadicSetContains(result, val) == true {
			result = append(result, val)
		}

	}
	return result
}
