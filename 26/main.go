package main

import (
	"fmt"
	"strings"
)

func checkIfUnique(s string) bool {
	lowS := strings.ToLower(s)
	runes := []rune(lowS)

	var runesMap = make(map[rune]int)

	for _, r := range runes {
		runesMap[r] += 1
	}

	for _, rm := range runesMap {
		if rm > 1 {
			return false
		}
	}
	return true
}

func main() {
	s := "abcd"
	fmt.Println(checkIfUnique(s)) // true

	s1 := "abCdefAaf"
	fmt.Println(checkIfUnique(s1)) // true

	s2 := "aabcd"
	fmt.Println(checkIfUnique(s2)) // false
}
