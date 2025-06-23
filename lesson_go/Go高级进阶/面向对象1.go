package main

import (
	"fmt"
	"math"
)

const pi = math.Pi

type Shape interface {
	Perimeter() float64
	Area() float64
}

// 题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
// 考察点 ：接口的定义与实现、面向对象编程风格。

// 长方形
type Rectangle struct {
	len   int
	width int
}

// 圆
type Circle struct {
	radius float64
}

// 实现接口方法
func (a *Rectangle) Perimeter() int {
	return (a.len + a.width) * 2
}
func (a *Rectangle) Area() int {
	return a.len * a.width
}

func (a *Circle) Perimeter() float64 {
	return float64(pi * a.radius * 2)
}
func (a *Circle) Area() float64 {
	return float64(pi * a.radius * a.radius)
}

func main() {
	rect := &Rectangle{len: 5, width: 3}
	circ := &Circle{radius: 2.5}
	fmt.Printf("矩形周长: %.2f\n", rect.Perimeter()) // 输出: 16.00
	fmt.Printf("矩形面积: %.2f\n", rect.Area())      // 输出: 15.00
	fmt.Printf("圆周长: %.2f\n", circ.Perimeter())  // 输出: 15.71
	fmt.Printf("圆面积: %.2f\n", circ.Area())       // 输出: 19.63
}
