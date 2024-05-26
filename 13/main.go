package main

import "fmt"

func swap(a, b int) (int, int) {
	a = a + b // сумма обоих чисел
	b = a - b // из общей суммы вычитаю b, получаю новое значение b = старому a
	a = a - b // так же из общей суммы вычитаю новое b(старое a), получаю a = старому b
	return a, b
}

func main() {
	num1 := 30
	num2 := 20
	fmt.Println("До swap:")
	fmt.Println("num1:", num1)
	fmt.Println("num2:", num2)

	num1, num2 = swap(num1, num2)

	fmt.Println("После swap:")
	fmt.Println("num1:", num1)
	fmt.Println("num2:", num2)
}
