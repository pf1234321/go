package main

import "fmt"

// sum 函数用于计算切片中所有元素的和，
// 并将结果发送到指定的整数通道中
func sum(s []int, c chan int) {
	sum := 0
	// 遍历切片中的每个元素并累加
	for _, v := range s {
		sum += v
	}
	c <- sum // 将计算结果发送到通道c
}

func main() {
	// 初始化一个整数切片
	s := []int{7, 2, 8, -9, 4, 0}

	// 创建一个无缓冲的整数通道
	c := make(chan int)

	// 启动第一个goroutine计算切片前半部分的和
	go sum(s[:len(s)/2], c)
	// 启动第二个goroutine计算切片后半部分的和
	go sum(s[len(s)/2:], c)

	// 从通道接收两个goroutine发送的结果
	// 注意：接收操作会阻塞直到两个结果都收到
	x, y := <-c, <-c

	// 输出两个部分的和以及它们的总和
	fmt.Println(x, y, x+y)
}
