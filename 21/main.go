package main

import "fmt"

// Представим что есть некий клиент, который ожидает реализацию методов этого интерфейса
type Foo interface {
	HelloFoo()
}

// Структура которая реализует интерфейс Foo
type FooImpl struct {
}

func (f *FooImpl) HelloFoo() {
	fmt.Println("i'm fooo")
}

func NewFoo() *FooImpl {
	return &FooImpl{}
}

// Так же представим что есть другой интерфейс с тем же функционалом (как у Foo)
type Boo interface {
	HelloBoo()
}

// Структура которая реализует интерфейс Boo
type BooImpl struct {
}

func (f *BooImpl) HelloBoo() {
	fmt.Println("i'm booo")
}

func NewBoo() *BooImpl {
	return &BooImpl{}
}

// нужно реализовать адаптер который будет использовать реализацию интерфейса Boo
// и переводитть на клиента который ожидает реализацию интерфейса Foo
type Adapter struct {
	boo BooImpl
}

func (a *Adapter) HelloFoo() {
	a.boo.HelloBoo()
}

func main() {
	adapter := &Adapter{boo: BooImpl{}}

	var foo Foo = adapter

	foo.HelloFoo()
}
