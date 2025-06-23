package main

import "fmt"

// fibonacci 函数使用 select 语句实现斐波那契数列的生成
// 通过 c 通道发送数列值，通过 quit 通道接收终止信号
func fibonacci(c, quit chan int) {
	x, y := 0, 1 // 初始化斐波那契数列的前两个值
	for {
		select {
		case c <- x: // 将当前值发送到数据通道
			x, y = y, x+y // 计算下一个斐波那契数
		case <-quit: // 从退出通道接收到信号时
			fmt.Println("quit") // 打印退出信息
			return              // 终止函数执行
		}
	}
}

// go func() { ... }()：启动一个匿名 goroutine，负责消费斐波那契数列。
// fibonacci(c, quit)：在主 goroutine 中调用，负责生成斐波那契数列。
// 通道 c 和 quit：实现两个 goroutine 之间的同步和通信。
func main() {
	c := make(chan int)    // 创建数据通道
	quit := make(chan int) // 创建退出信号通道
	// 启动一个匿名 goroutine 负责接收斐波那契数并控制退出
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("======================", <-c) // 从数据通道接收并打印斐波那契数
		}
		quit <- 0 // 发送退出信号到 quit 通道
	}()
	// 启动斐波那契数列生成器
	fibonacci(c, quit)
}
