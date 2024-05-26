package main

import (
	"fmt"
)

// Создаю структуру Human
type Human struct {
	Age  int
	Name string
}

// Добавляю метод getHuman для структуры Human
func (h *Human) getHuman() string {
	return fmt.Sprintf("Age is: %d, Name is: %s", h.Age, h.Name)
}

// Встраиваю структуру Human
type Action struct {
	Human Human
}

func main() {
	// создаю экземпляр структуры Action
	a := Action{Human{Age: 10, Name: "Artem"}}

	// Вызываю метод getHuman структуры Human у структуры Action

	s := a.Human.getHuman()
	fmt.Println(s)
}
