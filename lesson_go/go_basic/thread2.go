package main

import "fmt"

func sum2(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	//，go 关键字用于启动一个轻量级线程（goroutine），使函数在独立的执行流程中异步运行。
	go sum2(s[:len(s)/2], c)
	go sum2(s[len(s)/2:], c)
	x, y := <-c, <-c // 从通道 c 中接收

	fmt.Println(x, y, x+y)
}
