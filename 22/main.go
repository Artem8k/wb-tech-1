package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(2)
	a.Exp(a, big.NewInt(20), nil) // a = 2^20

	b := big.NewInt(3)
	b.Exp(b, big.NewInt(20), nil) // b = 3^20

	// Умножение
	mul := new(big.Int).Mul(a, b)
	fmt.Printf("Умножение: %s\n", mul.String())

	// Деление
	div := new(big.Int).Div(mul, a)
	fmt.Printf("Деление: %s\n", div.String())

	// Сложение
	sum := new(big.Int).Add(a, b)
	fmt.Printf("Сложение: %s\n", sum.String())

	// Вычитание
	sub := new(big.Int).Sub(a, b)
	fmt.Printf("Вычитание: %s\n", sub.String())
}
