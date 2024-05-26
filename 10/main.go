package main

import "fmt"

func main() {
	temperature := [...]float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	var temperatureMap = make(map[int][]float32)

	/* создаю ключ для группировки значений
	по этому ключу добавляю значения в слайс */
	for _, t := range temperature {
		key := int(t/10) * 10
		temperatureMap[key] = append(temperatureMap[key], t)
	}

	fmt.Println(temperatureMap)
}
