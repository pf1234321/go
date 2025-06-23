package main

import "fmt"

func main() {
	ch := make(chan int) // 创建无缓冲通道

	// 生成器协程：发送1到10的整数
	go func() {
		defer close(ch) // 发送完毕后关闭通道
		for i := 1; i <= 10; i++ {
			ch <- i // 发送整数到通道
		}
	}()

	// 接收器协程：从通道接收并打印
	go func() {
		for num := range ch { // 自动循环直到通道关闭
			fmt.Println(num)
		}
	}()

	// 主协程等待（实际项目中建议使用 sync.WaitGroup）
	fmt.Scanln() // 阻塞主协程，按回车键退出
}
