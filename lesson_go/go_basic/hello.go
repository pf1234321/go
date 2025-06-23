package main

// import "fmt"
// func main() {
// 	fmt.Println("Hello, World!")
// }

import "fmt"

func main() {
	var intVal int
	//_intVal = 1 // 这时候会产生编译错误，因为 intVal 已经声明，不需要重新声明
	intVal = 1 // 此时不会产生编译错误，因为有声明新的变量，因为 := 是一个声明语句
	fmt.Println(intVal)
	var stockcode int = 123
	var enddate string = "2020-12-31"
	var url string = "Code=%d&endDate=%s"
	var target_url string = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(target_url)
}
