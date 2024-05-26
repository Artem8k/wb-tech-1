package main

import (
	"fmt"
	"strconv"
)

// Проверяю i-й бит, является ли он 1 или 0, возвращаю true если 1 и false если 0
func checkBit(num int64, pos int64) bool {
	v := num & (1 << pos)

	return v > 0
}

func setBit(num int64, pos int64) int64 {

	if checkBit(num, pos) == true {
		num &^= 1 << pos //сбрасываю бит на 0
	} else {
		num |= 1 << pos // устанавливаю 1 на нужную позицию
	}
	return num
}

func main() {

	var num int64 = 30

	var position int64 = 1

	fmt.Println("Исходная последовательность битов", strconv.FormatInt(num, 2))

	changedNum := setBit(num, position)

	fmt.Println("Последовательность битов после функции setBit", strconv.FormatInt(changedNum, 2))
}
