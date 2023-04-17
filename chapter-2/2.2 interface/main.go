package main

import "fmt"

// Animal 定义接口
type Animal interface {
	Speak() string
}

// Cat 定义结构体
type Cat struct{}

func (c Cat) Speak() string {
	return "喵喵"
}

type Dog struct{}

func (d Dog) Speak() string {
	return "汪汪"
}

type Dog2 struct {
}

func main() {
	var animal Animal

	animal = Cat{}
	fmt.Println(animal.Speak()) // 输出“喵喵”

	animal = Dog{}
	fmt.Println(animal.Speak()) // 输出“汪汪”

	fmt.Println(animal.Speak()) // 输出“汪汪”

}
