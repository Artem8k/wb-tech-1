package main

import (
	"fmt"
	"strings"
)

func reverseString(s string) string {
	stringSlice := strings.Split(s, " ")
	var result string

	// начинаю отсчет цикла с конца слайса со строками
	// присваиваю результирующей переменной строки с конца слайса
	for i := len(stringSlice) - 1; i >= 0; i-- {
		elem := stringSlice[i]
		result += elem + " "
	}
	return result
}

func main() {
	dummyString := "snow dog sun"

	fmt.Println(reverseString(dummyString))
}
