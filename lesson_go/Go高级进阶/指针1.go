package main

import (
	"fmt"
)

func sayHello(point *int) {
	var temp int
	temp = *point
	temp = temp + 10
	*point = temp
}

func main() {
	var a int = 20 /* 声明实际变量 */
	sayHello(&a)   // 启动 Goroutine
	fmt.Println(a)
}
