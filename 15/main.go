package main

import (
	"fmt"
	"strings"
)

// var justString string

// закоменнтил глобальную переменную
// из функции сразу возвращаю слайс строки, без дополнительного копирования в глоб. переменную
func someFunc() string {
	v := createHugeString(1 << 10)
	return v[:100]
}

func createHugeString(size int) string {
	return strings.Repeat("a", size)
}

func main() {
	someFunc()
	fmt.Println("justString:", someFunc())
}
