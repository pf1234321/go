package main

import "fmt"

// 父结构体
type Animal struct {
	Name string
}

func Speak2(a *Animal) {
	fmt.Println(a.Name, "says hello!")
}

// 父结构体的方法
func (a *Animal) Speak() {
	fmt.Println(a.Name, "says hello!")
}

// 子结构体
type Dog struct {
	Animal // 嵌入 Animal 结构体
	Breed  string
}

func main() {
	dog := Dog{
		Animal: Animal{Name: "Buddy"},
		Breed:  "Golden Retriever",
	}

	dog.Speak() // 调用父结构体的方法
	fmt.Println("Breed:", dog.Breed)
}
