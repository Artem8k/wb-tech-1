package main

import "fmt"

func findIntersection(dummyStrings []string) {

	var stringsMap = make(map[string]int)
	// собираю множество с ключем: строка, и значением: кол-во повторений в последовательности
	for _, elem := range dummyStrings {
		stringsMap[elem] += 1
	}

	fmt.Println(stringsMap)
}

func main() {
	dummyStrings := []string{"cat", "cat", "dog", "cat", "tree"}
	findIntersection(dummyStrings)
}
