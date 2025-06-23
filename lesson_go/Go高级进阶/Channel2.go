package main

import (
	"fmt"
)

func fibonacci(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func main() {
	// 这里我们定义了一个可以存储整数类型的带缓冲通道
	// 缓冲区大小为10
	ch := make(chan int, 100)
	go fibonacci(ch)
	// range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
	// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
	// 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
	// 会结束，从而在接收第 11 个数据的时候就阻塞了。
	for i := range ch {
		fmt.Println(i)
	}
}
