package main

import "fmt"

func main() {
	// 定义一个匿名函数并将其赋值给变量add
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println("Anonymous function:", "1 + 2 =", add(1, 2))
	// 调用匿名函数
	result := add(3, 5)
	fmt.Println("3 + 5 =", result)

	// 在函数内部使用匿名函数
	multiply := func(x, y int) int {
		return x * y
	}

	product := multiply(4, 6)
	fmt.Println("4 * 6 =", product)

	// 将匿名函数作为参数传递给其他函数
	calculate := func(operation func(int, int) int, x, y int) int {
		return operation(x, y)
	}

	sum := calculate(add, 2, 8)
	fmt.Println("2 + 8 =", sum)

	// 也可以直接在函数调用中定义匿名函数
	difference := calculate(func(a, b int) int {
		return a - b
	}, 10, 4)
	fmt.Println("10 - 4 =", difference)
}
