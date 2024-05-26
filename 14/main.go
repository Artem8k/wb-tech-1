package main

import "fmt"

func checkType(variable interface{}) {

	switch variable.(type) {
	case int:
		fmt.Println("переменная типа int")
	case string:
		fmt.Println("переменная типа string")
	case bool:
		fmt.Println("переменная типа bool")
	case chan int:
		fmt.Println("переменная типа bidirectional channel int")
	}

}

func main() {
	num := 1
	checkType(num)

	str := "1"
	checkType(str)

	bol := true
	checkType(bol)

	ch := make(chan int)
	checkType(ch)
}
